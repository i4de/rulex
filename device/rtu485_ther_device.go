package device

import (
	"context"
	golog "log"
	"os"
	"sync"
	"time"

	"github.com/i4de/rulex/common"
	"github.com/i4de/rulex/core"
	"github.com/i4de/rulex/driver"
	"github.com/i4de/rulex/glogger"
	"github.com/i4de/rulex/typex"
	"github.com/i4de/rulex/utils"

	"github.com/goburrow/modbus"
	"github.com/mitchellh/mapstructure"
)

type rtu485_ther struct {
	typex.XStatus
	status     typex.DeviceState
	RuleEngine typex.RuleX
	driver     typex.XExternalDriver
	rtuHandler *modbus.RTUClientHandler
	mainConfig common.ModBusConfig
	rtuConfig  common.RTUConfig
	locker     sync.Locker
}

// Example: 0x02 0x92 0xFF 0x98
type __sensor_data struct {
	TEMP float32 `json:"temp"` //系数: 0.1
	HUM  float32 `json:"hum"`  //系数: 0.1
}

/*
*
* 温湿度传感器
*
 */
func NewRtu485Ther(e typex.RuleX) typex.XDevice {
	ther := new(rtu485_ther)
	ther.RuleEngine = e
	return ther
}

//  初始化
func (ther *rtu485_ther) Init(devId string, configMap map[string]interface{}) error {
	ther.PointId = devId
	if err := utils.BindSourceConfig(configMap, &ther.mainConfig); err != nil {
		return err
	}
	if errs := mapstructure.Decode(ther.mainConfig.Config, &ther.rtuConfig); errs != nil {
		glogger.GLogger.Error(errs)
		return errs
	}
	return nil
}

// 启动
func (ther *rtu485_ther) Start(cctx typex.CCTX) error {
	ther.Ctx = cctx.Ctx
	ther.CancelCTX = cctx.CancelCTX
	//
	// 串口配置固定写法
	ther.rtuHandler = modbus.NewRTUClientHandler(ther.rtuConfig.Uart)
	ther.rtuHandler.BaudRate = ther.rtuConfig.BaudRate
	ther.rtuHandler.DataBits = ther.rtuConfig.DataBits
	ther.rtuHandler.Parity = ther.rtuConfig.Parity
	ther.rtuHandler.StopBits = ther.rtuConfig.StopBits
	ther.rtuHandler.Timeout = time.Duration(ther.mainConfig.Frequency) * time.Second
	if core.GlobalConfig.AppDebugMode {
		ther.rtuHandler.Logger = golog.New(os.Stdout, "485-TEMP-HUMI-DEVICE: ", golog.LstdFlags)
	}
	if err := ther.rtuHandler.Connect(); err != nil {
		return err
	}
	client := modbus.NewClient(ther.rtuHandler)
	ther.driver = driver.NewRtu485THerDriver(ther.Details(),
		ther.RuleEngine, ther.mainConfig.Registers, ther.rtuHandler, client)
	//---------------------------------------------------------------------------------
	// Start
	//---------------------------------------------------------------------------------
	ther.status = typex.DEV_RUNNING
	go func(ctx context.Context, Driver typex.XExternalDriver) {
		ticker := time.NewTicker(time.Duration(ther.mainConfig.Frequency) * time.Second)
		defer ticker.Stop()
		buffer := make([]byte, common.T_64KB)
		for {
			<-ticker.C
			select {
			case <-ctx.Done():
				{
					ther.status = typex.DEV_STOP
					return
				}
			default:
				{
				}
			}
			ther.locker.Lock()
			n, err := Driver.Read(buffer)
			ther.locker.Unlock()
			if err != nil {
				glogger.GLogger.Error(err)
			} else {
				ther.RuleEngine.WorkDevice(ther.Details(), string(buffer[:n]))
			}
		}

	}(ther.Ctx, ther.driver)
	return nil
}

// 从设备里面读数据出来
func (ther *rtu485_ther) OnRead(data []byte) (int, error) {

	n, err := ther.driver.Read(data)
	if err != nil {
		glogger.GLogger.Error(err)
		ther.status = typex.DEV_STOP
	}
	return n, err
}

// 把数据写入设备
func (ther *rtu485_ther) OnWrite(_ []byte) (int, error) {
	return 0, nil
}

// 设备当前状态
func (ther *rtu485_ther) Status() typex.DeviceState {
	return typex.DEV_RUNNING
}

// 停止设备
func (ther *rtu485_ther) Stop() {
	if ther.rtuHandler != nil {
		ther.rtuHandler.Close()
	}
	ther.CancelCTX()
}

// 设备属性，是一系列属性描述
func (ther *rtu485_ther) Property() []typex.DeviceProperty {
	return []typex.DeviceProperty{}
}

// 真实设备
func (ther *rtu485_ther) Details() *typex.Device {
	return ther.RuleEngine.GetDevice(ther.PointId)
}

// 状态
func (ther *rtu485_ther) SetState(status typex.DeviceState) {
	ther.status = status

}

// 驱动
func (ther *rtu485_ther) Driver() typex.XExternalDriver {
	return ther.driver
}
