package grpc

import (
	"fmt"
	"net"

	hexpb "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc/pb"
	appPorts "github.com/AksAman/hexarch/internal/ports/app"
	"github.com/AksAman/hexarch/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	logger = utils.InitializeLogger("adapters.frameworks.left.grpc")
}

type Adapter struct {
	api      appPorts.APIPort
	listener net.Listener
}

// constructor
func NewAdapter(api appPorts.APIPort, addr string) (*Adapter, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener: %v", err)
	}
	return &Adapter{
		api:      api,
		listener: listener,
	}, nil
}

func NewAdapterWithListener(api appPorts.APIPort, listener net.Listener) *Adapter {
	return &Adapter{
		api:      api,
		listener: listener,
	}
}

func (rpcAdapter *Adapter) Run() {
	arithmetcServiceServer := rpcAdapter
	grpcServer := grpc.NewServer()
	hexpb.RegisterArithmeticServiceServer(
		grpcServer,
		arithmetcServiceServer,
	)
	logger.Infof("gRPC server starting on port: %v", rpcAdapter.listener.Addr())
	if err := grpcServer.Serve(rpcAdapter.listener); err != nil {
		logger.Fatalf("server failed to serve gRPCServer over %v: %v", rpcAdapter.listener.Addr(), err)
	}
}
