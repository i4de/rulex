package source

import (
	"context"
	"fmt"
	"net"

	"github.com/i4de/rulex/common"
	"github.com/i4de/rulex/core"
	"github.com/i4de/rulex/glogger"
	"github.com/i4de/rulex/rulexrpc"
	"github.com/i4de/rulex/typex"
	"github.com/i4de/rulex/utils"

	"google.golang.org/grpc"
)

const (
	DefaultTransport = "tcp"
)

type RulexRpcServer struct {
	grpcInEndSource *grpcInEndSource
	rulexrpc.UnimplementedRulexRpcServer
}

//
// Source interface
//
type grpcInEndSource struct {
	typex.XStatus
	rulexServer *RulexRpcServer
	rpcServer   *grpc.Server
	mainConfig  common.GrpcConfig
}

//
func NewGrpcInEndSource(e typex.RuleX) typex.XSource {
	g := grpcInEndSource{}
	g.RuleEngine = e
	g.mainConfig = common.GrpcConfig{}
	return &g
}

/*
*
* Init
*
 */
func (g *grpcInEndSource) Init(inEndId string, configMap map[string]interface{}) error {
	g.PointId = inEndId
	if err := utils.BindSourceConfig(configMap, &g.mainConfig); err != nil {
		return err
	}
	return nil
}

//
func (g *grpcInEndSource) Start(cctx typex.CCTX) error {
	g.Ctx = cctx.Ctx
	g.CancelCTX = cctx.CancelCTX

	listener, err := net.Listen(DefaultTransport, fmt.Sprintf(":%d", g.mainConfig.Port))
	if err != nil {
		return err
	}
	// Important !!!
	g.rpcServer = grpc.NewServer()
	g.rulexServer = new(RulexRpcServer)
	g.rulexServer.grpcInEndSource = g
	//
	rulexrpc.RegisterRulexRpcServer(g.rpcServer, g.rulexServer)

	go func(c context.Context) {
		glogger.GLogger.Info("RulexRpc source started on", listener.Addr())
		g.rpcServer.Serve(listener)
	}(g.Ctx)

	return nil
}

//
func (g *grpcInEndSource) DataModels() []typex.XDataModel {
	return g.XDataModels
}

//
func (g *grpcInEndSource) Stop() {
	if g.rpcServer != nil {
		g.rpcServer.Stop()
	}
	g.CancelCTX()

}
func (g *grpcInEndSource) Reload() {

}
func (g *grpcInEndSource) Pause() {

}
func (g *grpcInEndSource) Status() typex.SourceState {
	return typex.SOURCE_UP
}

func (g *grpcInEndSource) Test(inEndId string) bool {
	return true
}

func (g *grpcInEndSource) Enabled() bool {
	return true
}

func (g *grpcInEndSource) Details() *typex.InEnd {
	return g.RuleEngine.GetInEnd(g.PointId)
}

func (*grpcInEndSource) Driver() typex.XExternalDriver {
	return nil
}
func (*grpcInEndSource) Configs() *typex.XConfig {
	return core.GenInConfig(typex.GRPC, "GRPC", common.GrpcConfig{})
}

//
func (r *RulexRpcServer) Work(ctx context.Context, in *rulexrpc.Data) (*rulexrpc.Response, error) {
	ok, err := r.grpcInEndSource.RuleEngine.WorkInEnd(
		r.grpcInEndSource.RuleEngine.GetInEnd(r.grpcInEndSource.PointId),
		in.Value,
	)
	if ok {
		return &rulexrpc.Response{
			Code:    0,
			Message: "OK",
		}, nil
	} else {
		return &rulexrpc.Response{
			Code:    1,
			Message: err.Error(),
		}, err
	}

}

//
// ??????
//
func (*grpcInEndSource) Topology() []typex.TopologyPoint {
	return []typex.TopologyPoint{}
}

//
// ?????????????????????
//
func (*grpcInEndSource) DownStream([]byte) {}

//
// ????????????
//
func (*grpcInEndSource) UpStream() {}
