package grpc

import (
	"context"

	hexpb "github.com/AksAman/hexarch/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (rpcAdapter *Adapter) GetAddition(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error) {
	ans := &hexpb.Answer{}

	// TODO: Check this
	// TODO: This is a temporary solution, as the generated pb.go files do not accept 0 (check this, may be wrong)
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "a and b must be non-zero")
	}

	answer, err := rpcAdapter.api.GetAddition(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &hexpb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (rpcAdapter *Adapter) GetSubtraction(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error) {
	ans := &hexpb.Answer{}

	// TODO: Check this
	// TODO: This is a temporary solution, as the generated pb.go files do not accept 0 (check this, may be wrong)
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "a and b must be non-zero")
	}

	answer, err := rpcAdapter.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &hexpb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (rpcAdapter *Adapter) GetMultiplication(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error) {
	ans := &hexpb.Answer{}

	// TODO: Check this
	// TODO: This is a temporary solution, as the generated pb.go files do not accept 0 (check this, may be wrong)
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "a and b must be non-zero")
	}

	answer, err := rpcAdapter.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &hexpb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (rpcAdapter *Adapter) GetDivision(ctx context.Context, req *hexpb.OperationParameters) (*hexpb.Answer, error) {
	ans := &hexpb.Answer{}

	// TODO: Check this
	// TODO: This is a temporary solution, as the generated pb.go files do not accept 0 (check this, may be wrong)
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "a and b must be non-zero")
	}

	answer, err := rpcAdapter.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &hexpb.Answer{
		Value: answer,
	}
	return ans, nil
}
