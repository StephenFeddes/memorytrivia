package createlobby

import (
	"context"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc/protobufs"
)

type CreateLobby struct{}

func NewCreateLobby() *CreateLobby {
	return &CreateLobby{}
}

func (c *CreateLobby) Execute(ctx context.Context, lobby *pb.CreateLobbyRequest) (string, error) {
	return "success", nil
}