// This will be the entry point for our application
// orchestrating the startup of the application and
// dependency injection etc.

package main

import (
	"github.com/AksAman/hexarch/config"
	"github.com/AksAman/hexarch/internal/adapters/app/api"
	"github.com/AksAman/hexarch/internal/adapters/core/arithmetic"
	gRPC "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc"
	"github.com/AksAman/hexarch/internal/adapters/framework/right/db"
	"github.com/AksAman/hexarch/internal/ports"
	"github.com/AksAman/hexarch/utils"
	"go.uber.org/zap"
)

var (
	logger    *zap.SugaredLogger
	appConfig *config.Config
)

// ports
var (
	// domain/core ports
	coreAdapter ports.ArithmeticPort
	// application ports
	appAdapter ports.APIPort
	// framework ports
	dbAdapter   ports.DBPort
	gRPCAdapter ports.GRPCPort
)

func init() {
	logger = utils.InitializeLogger("cmd.main")

	var err error
	appConfig, err = config.LoadConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
}

func main() {
	coreAdapter = createCoreAdapter()
	dbAdapter = createDBAdapter()
	appAdapter = createAppAdapter(coreAdapter, dbAdapter)
	gRPCAdapter = createGRPCAdapter(appAdapter)

	gRPCAdapter.Run()
}

func createCoreAdapter() ports.ArithmeticPort {
	return arithmetic.NewAdapter()
}

func createAppAdapter(arithPort ports.ArithmeticPort, dbPort ports.DBPort) ports.APIPort {
	return api.NewAdapter(arithPort, dbPort)
}

func createDBAdapter() ports.DBPort {
	driverName := "postgres"
	dataSourceName := appConfig.GetPGConnectionString()
	logger.Debugf("dataSourceName: %q", dataSourceName)
	logger.Debugf("driverName: %q", driverName)

	var err error
	dbPort, err := db.NewAdapter(driverName, dataSourceName)
	if err != nil {
		logger.Fatalf("failed to create db adapter: %v", err)
	}
	defer dbPort.CloseDBConnection()
	return dbPort
}

func createGRPCAdapter(apiPort ports.APIPort) ports.GRPCPort {
	return gRPC.NewAdapter(apiPort)
}
