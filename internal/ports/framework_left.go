package ports

import (
	"context"

	hexpb "github.com/AksAman/hexarch/internal/adpaters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
	GetSubtraction(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
	GetMultiplication(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
	GetDivision(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error)
}
