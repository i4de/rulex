package source

import (
	"context"
	"fmt"
	"net/http"

	"github.com/i4de/rulex/common"
	"github.com/i4de/rulex/core"
	"github.com/i4de/rulex/glogger"
	"github.com/i4de/rulex/typex"
	"github.com/i4de/rulex/utils"

	"github.com/gin-gonic/gin"
)

//
type httpInEndSource struct {
	typex.XStatus
	engine     *gin.Engine
	mainConfig common.HostConfig
}

func NewHttpInEndSource(e typex.RuleX) typex.XSource {
	h := httpInEndSource{}
	gin.SetMode(gin.ReleaseMode)
	h.engine = gin.New()
	h.RuleEngine = e
	return &h
}
func (*httpInEndSource) Configs() *typex.XConfig {
	return core.GenInConfig(typex.HTTP, "HTTP", common.HTTPConfig{})
}
func (hh *httpInEndSource) Init(inEndId string, configMap map[string]interface{}) error {
	hh.PointId = inEndId
	if err := utils.BindSourceConfig(configMap, &hh.mainConfig); err != nil {
		return err
	}
	return nil
}

//
func (hh *httpInEndSource) Start(cctx typex.CCTX) error {
	hh.Ctx = cctx.Ctx
	hh.CancelCTX = cctx.CancelCTX

	hh.engine.POST("/in", func(c *gin.Context) {
		type Form struct {
			Data string
		}
		var inForm Form
		err := c.BindJSON(&inForm)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			hh.RuleEngine.WorkInEnd(hh.RuleEngine.GetInEnd(hh.PointId), inForm.Data)
			c.JSON(200, gin.H{
				"message": "ok",
				"data":    inForm,
			})
		}
	})

	go func(ctx context.Context) {
		err := http.ListenAndServe(fmt.Sprintf(":%v", hh.mainConfig), hh.engine)
		if err != nil {
			glogger.GLogger.Error(err)
			return
		}
	}(hh.Ctx)
	glogger.GLogger.Info("HTTP source started on" + " [0.0.0.0]:" + fmt.Sprintf("%v", hh.mainConfig.Port))

	return nil
}

//
func (mm *httpInEndSource) DataModels() []typex.XDataModel {
	return mm.XDataModels
}

//
func (hh *httpInEndSource) Stop() {
	hh.CancelCTX()

}
func (hh *httpInEndSource) Reload() {

}
func (hh *httpInEndSource) Pause() {

}
func (hh *httpInEndSource) Status() typex.SourceState {
	return typex.SOURCE_UP
}

func (hh *httpInEndSource) Test(inEndId string) bool {
	return true
}

func (hh *httpInEndSource) Enabled() bool {
	return hh.Enable
}
func (hh *httpInEndSource) Details() *typex.InEnd {
	return hh.RuleEngine.GetInEnd(hh.PointId)
}

func (*httpInEndSource) Driver() typex.XExternalDriver {
	return nil
}

//
// ??????
//
func (*httpInEndSource) Topology() []typex.TopologyPoint {
	return []typex.TopologyPoint{}
}

//
// ?????????????????????
//
func (*httpInEndSource) DownStream([]byte) {}

//
// ????????????
//
func (*httpInEndSource) UpStream() {}
