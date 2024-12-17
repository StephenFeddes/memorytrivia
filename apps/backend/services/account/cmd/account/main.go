package main

import (
	"log"
	"os"

	"github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/database"
	gRPCServer "github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/grpc"
	httpServer "github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/http"
)


func main() {
	db, err := database.NewPostgresDB(os.Getenv("POSTGRES_URL")); if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.RunMigrations(); if err != nil {
		log.Fatal(err)
	}

	httpServer := httpServer.NewHTTPServer(":80", db)
	go httpServer.Run()

	grpcServer := gRPCServer.NewGRPCServer(":50051", db)
	grpcServer.Run()
}