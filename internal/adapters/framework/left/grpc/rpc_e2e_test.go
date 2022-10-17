package grpc_test

import (
	"context"
	"net"
	"testing"

	"github.com/AksAman/hexarch/config"
	"github.com/AksAman/hexarch/internal/adapters/app/api"
	"github.com/AksAman/hexarch/internal/adapters/core/arithmetic"
	hexgRPC "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc"
	hexpb "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc/pb"
	"github.com/AksAman/hexarch/internal/adapters/framework/right/pgdb"
	appPorts "github.com/AksAman/hexarch/internal/ports/app"
	corePorts "github.com/AksAman/hexarch/internal/ports/core"
	leftFrameworkPorts "github.com/AksAman/hexarch/internal/ports/framework/left"
	rightFrameworkPorts "github.com/AksAman/hexarch/internal/ports/framework/right"
	"github.com/AksAman/hexarch/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const buffSize = 1024 * 1024

var (
	logger       *zap.SugaredLogger
	buffListener *bufconn.Listener
	appConfig    *config.Config
)

func init() {
	logger = utils.InitializeLogger("cmd.main")
	var err error
	appConfig, err = config.LoadConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	// ports
	coreAdapter := createCoreAdapter()
	dbAdapter := createPGDBAdapter()
	// dbAdapter = createJSONDBAdapter()
	appAdapter := createAppAdapter(coreAdapter, dbAdapter)
	gRPCAdapter := createGRPCAdapter(appAdapter)

	go func() {
		gRPCAdapter.Run()
	}()
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

func createGRPCAdapter(apiPort appPorts.APIPort) leftFrameworkPorts.GRPCPort {
	buffListener = bufconn.Listen(buffSize)
	return hexgRPC.NewAdapterWithListener(apiPort, buffListener)
}

func createBufDialer(context.Context, string) (net.Conn, error) {
	return buffListener.Dial()
}

func getGRPConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(createBufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestAdapter_GetAddition(t *testing.T) {
	ctx := context.Background()
	grpcClient := hexpb.NewArithmeticServiceClient(getGRPConnection(ctx, t))

	tests := []struct {
		name    string
		client  hexpb.ArithmeticServiceClient
		args    *hexpb.OperationParameters
		want    int32
		wantErr bool
	}{
		{name: "case1", client: grpcClient, args: &hexpb.OperationParameters{A: 5, B: 2}, want: int32(7), wantErr: false},
		{name: "zeroinputa", client: grpcClient, args: &hexpb.OperationParameters{A: 0, B: 2}, want: int32(0), wantErr: true},
		{name: "zeroinputb", client: grpcClient, args: &hexpb.OperationParameters{A: 2, B: 0}, want: int32(0), wantErr: true},
	}

	assert.Greater(t, len(tests), 0, "no test cases found")

	for _, testcase := range tests {
		t.Run(
			testcase.name,
			func(t *testing.T) {
				got, err := testcase.client.GetAddition(ctx, testcase.args)
				if testcase.wantErr {
					assert.NotNil(t, err, "expected error")
				} else {
					assert.Nil(t, err, "unexpected error")
					assert.Equal(t, testcase.want, got.Value, "unexpected result")
				}
			},
		)
	}
}

func TestAdapter_GetSubtraction(t *testing.T) {
	ctx := context.Background()
	grpcClient := hexpb.NewArithmeticServiceClient(getGRPConnection(ctx, t))

	tests := []struct {
		name    string
		client  hexpb.ArithmeticServiceClient
		args    *hexpb.OperationParameters
		want    int32
		wantErr bool
	}{
		{name: "positive subtraction", client: grpcClient, args: &hexpb.OperationParameters{A: 5, B: 2}, want: int32(5 - 2), wantErr: false},
		{name: "negative subtraction", client: grpcClient, args: &hexpb.OperationParameters{A: 2, B: 5}, want: int32(2 - 5), wantErr: false},
		{name: "zeroinputa", client: grpcClient, args: &hexpb.OperationParameters{A: 0, B: 2}, want: int32(0), wantErr: true},
		{name: "zeroinputb", client: grpcClient, args: &hexpb.OperationParameters{A: 2, B: 0}, want: int32(0), wantErr: true},
	}

	assert.Greater(t, len(tests), 0, "no test cases found")

	for _, testcase := range tests {
		t.Run(
			testcase.name,
			func(t *testing.T) {
				got, err := testcase.client.GetSubtraction(ctx, testcase.args)
				if testcase.wantErr {
					assert.NotNil(t, err, "expected error")
				} else {
					assert.Nil(t, err, "unexpected error")
					assert.Equal(t, testcase.want, got.Value, "unexpected result")
				}
			},
		)
	}
}

func TestAdapter_GetMultiplication(t *testing.T) {
	ctx := context.Background()
	grpcClient := hexpb.NewArithmeticServiceClient(getGRPConnection(ctx, t))

	tests := []struct {
		name    string
		client  hexpb.ArithmeticServiceClient
		args    *hexpb.OperationParameters
		want    int32
		wantErr bool
	}{
		{name: "case1", client: grpcClient, args: &hexpb.OperationParameters{A: 5, B: 2}, want: int32(5 * 2), wantErr: false},
		{name: "zeroinputa", client: grpcClient, args: &hexpb.OperationParameters{A: 0, B: 2}, want: int32(0), wantErr: true},
		{name: "zeroinputb", client: grpcClient, args: &hexpb.OperationParameters{A: 2, B: 0}, want: int32(0), wantErr: true},
	}

	assert.Greater(t, len(tests), 0, "no test cases found")

	for _, testcase := range tests {
		t.Run(
			testcase.name,
			func(t *testing.T) {
				got, err := testcase.client.GetMultiplication(ctx, testcase.args)
				if testcase.wantErr {
					assert.NotNil(t, err, "expected error")
				} else {
					assert.Nil(t, err, "unexpected error")
					assert.Equal(t, testcase.want, got.Value, "unexpected result")
				}
			},
		)
	}
}

func TestAdapter_GetDivision(t *testing.T) {
	ctx := context.Background()
	grpcClient := hexpb.NewArithmeticServiceClient(getGRPConnection(ctx, t))

	tests := []struct {
		name    string
		client  hexpb.ArithmeticServiceClient
		args    *hexpb.OperationParameters
		want    int32
		wantErr bool
	}{
		{name: "integer division", client: grpcClient, args: &hexpb.OperationParameters{A: 6, B: 2}, want: int32(3), wantErr: false},
		{name: "float division", client: grpcClient, args: &hexpb.OperationParameters{A: 5, B: 2}, want: int32(2), wantErr: false},
		{name: "division by 0", client: grpcClient, args: &hexpb.OperationParameters{A: 6, B: 0}, want: int32(0), wantErr: true},
	}

	assert.Greater(t, len(tests), 0, "no test cases found")

	for _, testcase := range tests {
		t.Run(
			testcase.name,
			func(t *testing.T) {
				got, err := testcase.client.GetDivision(ctx, testcase.args)
				if testcase.wantErr {
					assert.NotNil(t, err, "expected error")
				} else {
					assert.Nil(t, err, "unexpected error")
					assert.Equal(t, testcase.want, got.Value, "unexpected result")
				}
			},
		)
	}
}
