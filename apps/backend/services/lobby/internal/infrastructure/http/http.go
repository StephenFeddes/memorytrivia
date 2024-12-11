package http

import (
	"log"
	"net/http"

	getLobby "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/application/queries/getlobby"
	httpHandler "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/presentation/http"
)

// TODO: Use HTTP to send server-sent events, such as a player joining a lobby.

type httpServer struct {
	address string
}

func NewHTTPServer(address string) *httpServer {
	return &httpServer{address: address}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	prefix := "/api/v1/"

	getLobby := getLobby.NewGetLobby()
	lobbyHandler := httpHandler.NewLobbyHTTPHandler(getLobby)
	lobbyHandler.RegisterLobbyServiceHTTPServer(router, prefix+"lobby")

	log.Println("🚀 HTTP server is running on port %v", s.address)

	return http.ListenAndServe(s.address, router)
}