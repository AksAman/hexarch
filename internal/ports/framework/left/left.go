package leftFrameworkPorts

import (
	"context"

	hexpb "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
	GetSubtraction(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
	GetMultiplication(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
	GetDivision(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
}
