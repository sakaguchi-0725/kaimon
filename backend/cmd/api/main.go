package main

import (
	"backend/core"
	"backend/infra/db"
	"backend/presentation/server"
	"backend/registry"
	"log"
)

func main() {
	cfg, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := db.New(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo, err := registry.NewRepository(db, cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	uc := registry.NewUseCase(repo)

	logger := core.NewLogger(cfg.Env)

	srv := server.New(8080)
	srv.MapRoutes(cfg.FrontendURL, logger, uc)
	srv.Run()
}
