package grpc

import (
	"context"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/grpc/protobufs"
	"google.golang.org/grpc"
)

type AccountGetter interface {
	Execute(context.Context, uint32) (*pb.Account, error)
}

type AccountGRPCHandler struct {
	lobbyGetter AccountGetter
	pb.UnimplementedAccountServiceServer
}

func NewAccountGRPCHandler(lobbyGetter AccountGetter) *AccountGRPCHandler {
	return &AccountGRPCHandler{
		lobbyGetter: lobbyGetter,
	}
}

func (h *AccountGRPCHandler) RegisterAccountServiceGRPCServer(grpcServer *grpc.Server) {
	pb.RegisterAccountServiceServer(grpcServer, h)
}

func (h *AccountGRPCHandler) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	lobby, err := h.lobbyGetter.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return lobby, nil
}
