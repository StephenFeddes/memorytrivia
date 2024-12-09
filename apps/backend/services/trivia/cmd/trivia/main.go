package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/trivia/api/grpc/protobufs"
	"google.golang.org/grpc"
)

type TriviaServiceServer struct {
	pb.UnimplementedTriviaServiceServer
}

// Example method implementation (replace with actual methods from your proto)
func (t *TriviaServiceServer) GetQuestion(ctx context.Context, req *pb.GetQuestionRequest) (*pb.Question, error) {
	return &pb.Question{
		Id: 1,
		Prompt: "What is the capital of France?",
		Choices: []string{"New York", "Paris", "London", "Berlin"},
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the AccountServiceServer with the gRPC server
	pb.RegisterTriviaServiceServer(grpcServer, &TriviaServiceServer{})

	log.Printf("🚀 gRPC server is running on port %s", port)

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}