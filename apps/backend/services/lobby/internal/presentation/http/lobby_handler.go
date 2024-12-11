package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc/protobufs"
)

// Interfaces for the lobby service structs.
type LobbyGetter interface {
	Execute(context.Context, uint64) (*pb.Lobby, error)
}
type LobbyCreator interface {
	Execute(context.Context, *pb.CreateLobbyRequest) (string, error)
}

// LobbyHTTPHandler represents the HTTP handler for the lobby service.
type LobbyHTTPHandler struct {
	lobbyGetter LobbyGetter
}

// NewLobbyHTTPHandler creates a new instance of LobbyHTTPHandler.
func NewLobbyHTTPHandler(lobbyGetter LobbyGetter) *LobbyHTTPHandler {
	handler := &LobbyHTTPHandler{
		lobbyGetter: lobbyGetter,
	}

	return handler
}

func (h *LobbyHTTPHandler) RegisterLobbyServiceHTTPServer(router *http.ServeMux, prefix string) {
	router.HandleFunc(fmt.Sprintf("GET %s/{id}", prefix), h.getLobby)
} 

func (h *LobbyHTTPHandler) getLobby(w http.ResponseWriter, r *http.Request) {
	log.Println("Testing")
}

