package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Interfaces for the lobby service structs.
type LobbyGetter interface {
	Execute(context.Context, uint32) (*pb.Lobby, error)
}
type LobbyCreator interface {
	Execute(context.Context, uint32, uint32) (string, error)
}

// LobbyHTTPHandler represents the HTTP handler for the lobby service.
type LobbyHTTPHandler struct {
	lobbyGetter  LobbyGetter
	lobbyCreator LobbyCreator
}

// NewLobbyHTTPHandler creates a new instance of LobbyHTTPHandler.
func NewLobbyHTTPHandler(lobbyGetter LobbyGetter, lobbyCreator LobbyCreator) *LobbyHTTPHandler {
	handler := &LobbyHTTPHandler{
		lobbyGetter:  lobbyGetter,
		lobbyCreator: lobbyCreator,
	}

	return handler
}

func (h *LobbyHTTPHandler) RegisterLobbyServiceHTTPServer(router *http.ServeMux, prefix string) {
	router.HandleFunc("GET "+prefix+"/{id}", h.getLobby)
	router.HandleFunc("POST "+prefix, h.createLobby)
}

func (h *LobbyHTTPHandler) getLobby(w http.ResponseWriter, r *http.Request) {
	log.Println("Testing!")
	lobbyId := r.PathValue("id")

	log.Printf("Getting lobby with id: %s", lobbyId)

	// Use grpc.NewClient to create the client connection
	conn, err := grpc.NewClient("account:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client using the connection
	client := pb.NewAccountServiceClient(conn)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetAccountRequest{Id: 1}
	res, err := client.GetAccount(ctx, req)
	if err != nil {
		log.Fatalf("could not get account: %v", err)
	}
	log.Printf("Account: %v", res)
}

type CreateLobbyRequestJSON struct {
	Size    uint32 `json:"size"`
	OwnerId uint32 `json:"ownerId"`
}

func (h *LobbyHTTPHandler) createLobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
	}
	var payload CreateLobbyRequestJSON
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	status, err := h.lobbyCreator.Execute(r.Context(), payload.Size, payload.OwnerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(status))

	log.Printf("Creating lobby with payload: Size=%d, OwnerId=%d", payload.Size, payload.OwnerId)
}
