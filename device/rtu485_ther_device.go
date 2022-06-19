package device

import (
	"context"
	"rulex/driver"
	"rulex/typex"
	"rulex/utils"
	"time"

	"github.com/goburrow/modbus"
	"github.com/mitchellh/mapstructure"
	"github.com/ngaut/log"
)

type rtu485_ther struct {
	typex.XStatus
	status     typex.DeviceState
	RuleEngine typex.RuleX
	driver     typex.XExternalDriver
	slaverIds  []byte
}

/*
*
* 温湿度传感器
*
 */
func NewRtu485Ther(deviceId string, e typex.RuleX) typex.XDevice {
	tss := new(rtu485_ther)
	tss.PointId = deviceId
	tss.RuleEngine = e
	return tss
}

//  初始化
func (tss *rtu485_ther) Init(devId string, config map[string]interface{}) error {

	return nil
}

// 启动
func (tss *rtu485_ther) Start(cctx typex.CCTX) error {
	tss.Ctx = cctx.Ctx
	tss.CancelCTX = cctx.CancelCTX
	config := tss.RuleEngine.GetDevice(tss.PointId).Config
	var mainConfig modBusConfig
	if err := utils.BindSourceConfig(config, &mainConfig); err != nil {
		return err
	}
	var rtuConfig rtuConfig
	if errs := mapstructure.Decode(mainConfig.Config, &rtuConfig); errs != nil {
		log.Error(errs)
		return errs
	}

	// 串口配置固定写法
	handler := modbus.NewRTUClientHandler(rtuConfig.Uart)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.Timeout = time.Duration(*mainConfig.Timeout) * time.Second
	// handler.Logger = golog.New(os.Stdout, "485THerSource: ", log.LstdFlags)
	if err := handler.Connect(); err != nil {
		return err
	}
	client := modbus.NewClient(handler)
	tss.driver = driver.NewRtu485_THer_Driver(tss.Details(), tss.RuleEngine, client)
	tss.slaverIds = append(tss.slaverIds, mainConfig.SlaverIds...)
	//---------------------------------------------------------------------------------
	// Start
	//---------------------------------------------------------------------------------
	tss.status = typex.DEV_RUNNING
	for _, slaverId := range tss.slaverIds {
		go func(ctx context.Context, slaverId byte, rtuDriver typex.XExternalDriver, handler *modbus.RTUClientHandler) {
			ticker := time.NewTicker(time.Duration(5) * time.Second)
			defer ticker.Stop()
			buffer := make([]byte, 32) //32字节数据
			for {
				<-ticker.C
				select {
				case <-ctx.Done():
					{
						tss.status = typex.DEV_STOP
						return
					}
				default:
					{
					}
				}
				handler.SlaveId = slaverId // 配置ID
				n, err := rtuDriver.Read(buffer)
				if err != nil {
					log.Error(err)
				} else {
					tss.RuleEngine.WorkDevice(tss.RuleEngine.GetDevice(tss.PointId), string(buffer[:n]))
				}
			}

		}(tss.Ctx, slaverId, tss.driver, handler)
	}
	return nil
}

// 从设备里面读数据出来
func (tss *rtu485_ther) OnRead(data []byte) (int, error) {

	n, err := tss.driver.Read(data)
	if err != nil {
		log.Error(err)
		tss.status = typex.DEV_STOP
	}
	return n, err
}

// 把数据写入设备
func (tss *rtu485_ther) OnWrite(_ []byte) (int, error) {
	return 0, nil
}

// 设备当前状态
func (tss *rtu485_ther) Status() typex.DeviceState {
	return typex.DEV_RUNNING
}

// 停止设备
func (tss *rtu485_ther) Stop() {
	tss.CancelCTX()
}

// 设备属性，是一系列属性描述
func (tss *rtu485_ther) Property() []typex.DeviceProperty {
	return []typex.DeviceProperty{}
}

// 真实设备
func (tss *rtu485_ther) Details() *typex.Device {
	return tss.RuleEngine.GetDevice(tss.PointId)
}

// 状态
func (tss *rtu485_ther) SetState(status typex.DeviceState) {
	tss.status = status

}

// 驱动
func (tss *rtu485_ther) Driver() typex.XExternalDriver {
	return tss.driver
}