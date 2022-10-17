package grpc

import (
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
	api appPorts.APIPort
	hexpb.UnimplementedArithmeticServiceServer
}

// constructor
func NewAdapter(api appPorts.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (rpcAdapter *Adapter) Run() {
	port := ":9000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalf("server failed to listen on port: %v with error: %v", port, err)
	}
	arithmetcServiceServer := rpcAdapter
	grpcServer := grpc.NewServer()
	hexpb.RegisterArithmeticServiceServer(
		grpcServer,
		arithmetcServiceServer,
	)

	logger.Infof("gRPC server starting on port: %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("server failed to serve gRPCServer over %v: %v", port, err)
	}
}
