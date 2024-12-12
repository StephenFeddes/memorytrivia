package http

import (
	"log"
	"net/http"

	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/database/repository"
	getLobby "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/application/queries/getlobby"
	createLobby "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/application/commands/createlobby"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/database"
	httpHandler "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/presentation/http"
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

	lobbyRepository := repository.NewPostgresLobbyRepository(s.db)

	createLobby := createLobby.NewCreateLobby(lobbyRepository)
	getLobby := getLobby.NewGetLobby()

	prefix := "/lobby/api/v1"
	lobbyHandler := httpHandler.NewLobbyHTTPHandler(getLobby, createLobby)
	lobbyHandler.RegisterLobbyServiceHTTPServer(router, prefix+"/lobbies")

	log.Printf("🚀 HTTP server is running on port %v", s.address)

	return http.ListenAndServe(s.address, router)
}