package http

import (
	"context"
	"log"
	"net/http"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/grpc/protobufs"
)

type AccountGetter interface {
	Execute(context.Context, uint32) (*pb.Account, error)
}

type AccountHTTPHandler struct {
	accountGetter  AccountGetter
}

// NewLobbyHTTPHandler creates a new instance of LobbyHTTPHandler.
func NewAccountHTTPHandler(accountGetter AccountGetter) *AccountHTTPHandler {
	handler := &AccountHTTPHandler{
		accountGetter:  accountGetter,
	}

	return handler
}

func (h *AccountHTTPHandler) RegisterLobbyServiceHTTPServer(router *http.ServeMux, prefix string) {
	router.HandleFunc("GET "+prefix+"/{id}", h.getAccount)
}

func (h *AccountHTTPHandler) getAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Testing!")
}
