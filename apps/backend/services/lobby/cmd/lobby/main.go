package main

import (
	"os"

	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/http"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc"
)


func main() {
	httpServer := http.NewHTTPServer(os.Getenv("HTTP_PORT"))
	go httpServer.Run()

	grpcServer := grpc.NewGRPCServer(os.Getenv("GRPC_PORT"))
	grpcServer.Run()
}