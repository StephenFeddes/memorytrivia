package grpc

import (
	"context"
	"fmt"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
	pb "github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/grpc/protobufs"
	"google.golang.org/grpc"
)

type AccountGetter interface {
	Execute(context.Context, uint32) (*entity.Account, error)
}

type AccountGRPCHandler struct {
	accountGetter AccountGetter
	pb.UnimplementedAccountServiceServer
}

func NewAccountGRPCHandler(accountGetter AccountGetter) *AccountGRPCHandler {
	return &AccountGRPCHandler{
		accountGetter: accountGetter,
	}
}

func (h *AccountGRPCHandler) RegisterAccountServiceGRPCServer(grpcServer *grpc.Server) {
	pb.RegisterAccountServiceServer(grpcServer, h)
}

func (h *AccountGRPCHandler) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	//account, err := h.accountGetter.Execute(ctx, req.Id)
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println("GetAccount")
	return &pb.Account{
		Id:       1,
		Username: "Stephen",
		Email:    "sfeddes@outlook.com",
	}, nil
}
