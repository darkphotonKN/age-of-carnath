package main

import (
	"fmt"
	"log"

	"github.com/darkphotonKN/age-of-carnath/internal/routes"
	"github.com/darkphotonKN/age-of-carnath/internal/server"
)

func main() {
	port := ":4111"
	server := server.NewMultiplayerServer(port)

	go server.MessageHub()

	// routes setup
	routes := routes.SetupRoutes(server)

	fmt.Printf("Server listening on port %s.\n", port)

	err := routes.Run(server.ListenAddr)

	if err != nil {
		log.Panic("Unable to start server. Err:", err)
	}
}
