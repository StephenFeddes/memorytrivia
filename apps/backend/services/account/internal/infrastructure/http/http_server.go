package http

import (
	"log"
	"net/http"
	"os"

	"github.com/StephenFeddes/memorytrivia/account/internal/application/command/loginaccount"
	"github.com/StephenFeddes/memorytrivia/account/internal/application/command/createaccount"
	"github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/auth"
	"github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/database"
	"github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/database/repository"
	handler "github.com/StephenFeddes/memorytrivia/account/internal/presentation/http"
)

// TODO: Use HTTP to send server-sent events, such as a player joining a lobby.

type httpServer struct {
	address string
	db      *database.PostgresDB
}

func NewHTTPServer(address string, db *database.PostgresDB) *httpServer {
	return &httpServer{address: address, db: db}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	accountRepository := repository.NewPostgresAccountRepository(s.db)
	jwtManager := auth.NewJWT(os.Getenv("JWT_SECRET"))

	createAccount := createaccount.NewCreateAccount(accountRepository, auth.NewBCrypt(), jwtManager)
	loginAccount := loginaccount.NewLoginAccount(accountRepository, auth.NewBCrypt(), jwtManager)

	prefix := "/account/api/v1"
	accountHandler := handler.NewAccountHTTPHandler(jwtManager, createAccount, loginAccount)
	accountHandler.RegisterAccountServiceHTTPServer(router, prefix+"/accounts")

	log.Printf("🚀 HTTP server is running on port %v", s.address)

	return http.ListenAndServe(s.address, router)
}
