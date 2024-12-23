package grpc

import (
	"log"
	"net"

	getLobby "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/application/queries/getlobby"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/database"
	grpcHandler "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/presentation/grpc"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	address string
	db 	*database.PostgresDB
}

func NewGRPCServer(address string, db *database.PostgresDB) *gRPCServer {
	return &gRPCServer{address: address, db: db}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	getLobby := getLobby.NewGetLobby()
	lobbyHandler := grpcHandler.NewLobbyGRPCHandler(getLobby)
	lobbyHandler.RegisterLobbyServiceGRPCServer(grpcServer)

	log.Printf("🚀 gRPC server is running on port %v", s.address)

	return grpcServer.Serve(lis)
}
