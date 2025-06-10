package main

import "backend/presentation/server"

func main() {
	srv := server.New(8080)
	srv.MapRoutes()
	srv.Run()
}
