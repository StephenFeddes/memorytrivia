package main

import (
	"log"
	"os"

	"github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/database"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/grpc"
	"github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/http"
)


func main() {
	db, err := database.NewPostgresDB(os.Getenv("POSTGRES_URL")); if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.RunMigrations(); if err != nil {
		log.Fatal(err)
	}

	httpServer := http.NewHTTPServer(":80", db)
	go httpServer.Run()

	grpcServer := grpc.NewGRPCServer(":50051", db)
	grpcServer.Run()
}