package grpc

import (
	"context"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc/protobufs"
	"google.golang.org/grpc"
)

type LobbyGetter interface {
	Execute(context.Context, uint64) (*protobufs.Lobby, error)
}

type LobbyGRPCHandler struct {
	lobbyGetter LobbyGetter
	protobufs.UnimplementedLobbyServiceServer
}

func NewLobbyGRPCHandler(lobbyGetter LobbyGetter) *LobbyGRPCHandler {
	return &LobbyGRPCHandler{
		lobbyGetter: lobbyGetter,
	}
}

func (h *LobbyGRPCHandler) RegisterLobbyServiceGRPCServer(grpcServer *grpc.Server) {
	protobufs.RegisterLobbyServiceServer(grpcServer, h)
}

func (h *LobbyGRPCHandler) GetLobby(ctx context.Context, req *protobufs.GetLobbyRequest) (*protobufs.Lobby, error) {
	lobby, err := h.lobbyGetter.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return lobby, nil
}
