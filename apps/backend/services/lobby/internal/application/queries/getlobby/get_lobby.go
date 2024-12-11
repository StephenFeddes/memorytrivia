package getlobby

import (
	"context"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc/protobufs"
)

type GetLobby struct {}

func NewGetLobby() *GetLobby {
	return &GetLobby{}
}

func (g *GetLobby) Execute(ctx context.Context, id uint64) (*pb.Lobby, error) {
	return &pb.Lobby{
		Id: 1,
		Size: 2,
		OwnerId: 1,
		MemberIds: []uint64{1, 2},
	}, nil
}
