// This will be the entry point for our application
// orchestrating the startup of the application and
// dependency injection etc.

package main

import (
	"fmt"

	"github.com/AksAman/hexarch/config"
	"github.com/AksAman/hexarch/internal/adapters/app/api"
	"github.com/AksAman/hexarch/internal/adapters/core/arithmetic"
	gRPC "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc"
	"github.com/AksAman/hexarch/internal/adapters/framework/right/jsondb"
	"github.com/AksAman/hexarch/internal/adapters/framework/right/pgdb"
	appPorts "github.com/AksAman/hexarch/internal/ports/app"
	corePorts "github.com/AksAman/hexarch/internal/ports/core"
	leftFrameworkPorts "github.com/AksAman/hexarch/internal/ports/framework/left"
	rightFrameworkPorts "github.com/AksAman/hexarch/internal/ports/framework/right"
	"github.com/AksAman/hexarch/utils"
	"go.uber.org/zap"
)

var (
	logger    *zap.SugaredLogger
	appConfig *config.Config
)

func init() {
	logger = utils.InitializeLogger("cmd.main")

	var err error
	appConfig, err = config.LoadConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
}

// ports
var (
	// domain/core ports
	coreAdapter corePorts.ArithmeticPort
	// application ports
	appAdapter appPorts.APIPort
	// framework ports
	dbAdapter   rightFrameworkPorts.DBPort
	gRPCAdapter leftFrameworkPorts.GRPCPort
)

func main() {
	coreAdapter = createCoreAdapter()
	dbAdapter = createPGDBAdapter()
	// dbAdapter = createJSONDBAdapter()
	defer dbAdapter.CloseDBConnection()
	appAdapter = createAppAdapter(coreAdapter, dbAdapter)
	gRPCAdapter = createGRPCAdapter(appAdapter)

	gRPCAdapter.Run()
}

func createCoreAdapter() corePorts.ArithmeticPort {
	return arithmetic.NewAdapter()
}

func createAppAdapter(arithPort corePorts.ArithmeticPort, dbPort rightFrameworkPorts.DBPort) appPorts.APIPort {
	return api.NewAdapter(arithPort, dbPort)
}

func createPGDBAdapter() rightFrameworkPorts.DBPort {
	driverName := "postgres"
	dataSourceName := appConfig.GetDBConnectionString()
	logger.Debugf("dataSourceName: %q", dataSourceName)
	logger.Debugf("driverName: %q", driverName)

	var err error
	dbPort, err := pgdb.NewAdapter(driverName, dataSourceName)
	if err != nil {
		logger.Fatalf("failed to create db adapter: %v", err)
	}
	return dbPort
}

func createJSONDBAdapter() rightFrameworkPorts.DBPort {
	var err error
	dbPort, err := jsondb.NewAdapter(appConfig.JSONDatabaseFilepath)
	if err != nil {
		logger.Fatalf("failed to create db adapter: %v", err)
	}
	return dbPort
}

func createGRPCAdapter(apiPort appPorts.APIPort) leftFrameworkPorts.GRPCPort {
	addr := fmt.Sprintf(":%d", appConfig.GRPCPort)
	adapter, err := gRPC.NewAdapter(apiPort, addr)
	if err != nil {
		logger.Fatalf("failed to create grpc adapter: %v", err)
	}
	return adapter
}
