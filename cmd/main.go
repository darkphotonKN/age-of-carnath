package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/darkphotonKN/age-of-carnath/config/db"
	"github.com/darkphotonKN/age-of-carnath/internal/routes"
	"github.com/darkphotonKN/age-of-carnath/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// env setup
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// db setup
	db := config.InitDB()
	defer db.Close()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := server.NewMultiplayerServer(port)
	// init messagehub concurrently
	go server.MessageHub()

	// routes setup
	routes := routes.SetupRoutes(server)

	fmt.Printf("Server listening on port %s.\n", port)

	err := routes.Run(server.ListenAddr)

	if err != nil {
		log.Panic("Unable to start server. Err:", err)
	}
}
