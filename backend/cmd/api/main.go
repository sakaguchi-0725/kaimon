package main

import (
	"backend/presentation/server"
	"backend/registry"
)

func main() {
	srv := server.New(8080)

	uc := registry.NewUseCase()
	srv.MapRoutes(uc)

	srv.Run()
}
