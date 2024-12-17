package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
)

type JWTValidator interface {
	ValidateToken(token string) (*entity.Account, error)
}
type AccountCreator interface {
	Execute(ctx context.Context, username string, email string, password string) (string, uint32, error)
}
type LoginHandler interface {
    Execute(ctx context.Context, username string, email string, password string) (string, uint32, error)
}

type AccountHTTPHandler struct {
	jwtValidator   JWTValidator
	accountCreator AccountCreator
    loginHandler   LoginHandler
}

// NewLobbyHTTPHandler creates a new instance of LobbyHTTPHandler.
func NewAccountHTTPHandler(j JWTValidator, a AccountCreator, l LoginHandler) *AccountHTTPHandler {
	handler := &AccountHTTPHandler{
		jwtValidator:   j,
		accountCreator: a,
        loginHandler:   l,
	}

	return handler
}

func (h *AccountHTTPHandler) RegisterAccountServiceHTTPServer(router *http.ServeMux, prefix string) {
	router.HandleFunc("POST "+prefix, h.createAccount)
    router.HandleFunc("POST "+prefix+"/tokens", h.loginAccount)
    router.HandleFunc("GET "+prefix+"/tokens", h.ValidateJWT)
}

func (h *AccountHTTPHandler) ValidateJWT(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	if header == "" {
		WriteError(w, http.StatusUnauthorized, errors.New("no token provided"))
		return
	}

	account, err := h.jwtValidator.ValidateToken(header)
	if err != nil {
		WriteError(w, http.StatusUnauthorized, err)
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{
		"account": map[string]any{
			"id":       account.ID,
			"username": account.Username,
			"email":    account.Email,
		},
	})
}

func (h *AccountHTTPHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ParseJSON(r, &payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
        return
	}

	jwt, accountID, err := h.accountCreator.Execute(r.Context(), payload.Username, payload.Email, payload.Password)
	if err != nil {
		WriteError(w, http.StatusConflict, err)
        return
	}

	WriteJSON(w, 201, map[string]any{
		"jwtToken": jwt,
		"account": map[string]any{
			"id":       accountID,
			"username": payload.Username,
			"email":    payload.Email,
		},
	})
}

func (h *AccountHTTPHandler) loginAccount(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ParseJSON(r, &payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
        return
	}

	jwt, accountID, err := h.loginHandler.Execute(r.Context(), payload.Username, payload.Email, payload.Password)
	if err != nil {
		WriteError(w, http.StatusUnauthorized, err)
        return
	}

	WriteJSON(w, 201, map[string]any{
		"jwtToken": jwt,
		"account": map[string]any{
			"id":       accountID,
			"username": payload.Username,
			"email":    payload.Email,
		},
	})
}
