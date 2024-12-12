package getaccount

import (
	"context"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/grpc/protobufs"
)

type GetAccount struct{}

func NewGetAccount() *GetAccount {
	return &GetAccount{}
}

func (g *GetAccount) Execute(ctx context.Context, id uint32) (*pb.Account, error) {
	return &pb.Account{
		Id:       1,
		Username: "test",
		Password: "testpassword",
		Email:    "testemail",
	}, nil
}
