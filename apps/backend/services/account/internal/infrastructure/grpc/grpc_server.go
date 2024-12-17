package grpc

import (
	"log"
	"net"

	getAccount "github.com/StephenFeddes/memorytrivia/account/internal/application/query/getaccount"
	"github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/database"
	handler "github.com/StephenFeddes/memorytrivia/account/internal/presentation/grpc"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	address string
	db      *database.PostgresDB
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

	getAccount := getAccount.NewGetAccount()
	accountHandler := handler.NewAccountGRPCHandler(getAccount)
	accountHandler.RegisterAccountServiceGRPCServer(grpcServer)

	log.Printf("🚀 gRPC server is running on port %v", s.address)

	return grpcServer.Serve(lis)
}
