package http

import (
	"log"
	"net/http"

	getAccount "github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/application/queries/getaccount"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/database"
	httpHandler "github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/presentation/http"
)

// TODO: Use HTTP to send server-sent events, such as a player joining a lobby.

type httpServer struct {
	address string
	db 	*database.PostgresDB
}

func NewHTTPServer(address string, db *database.PostgresDB) *httpServer {
	return &httpServer{address: address, db: db}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	getAccount := getAccount.NewGetAccount()

	prefix := "/lobby/api/v1"
	lobbyHandler := httpHandler.NewAccountHTTPHandler(getAccount)
	lobbyHandler.RegisterLobbyServiceHTTPServer(router, prefix+"/lobbies")

	log.Printf("🚀 HTTP server is running on port %v", s.address)

	return http.ListenAndServe(s.address, router)
}