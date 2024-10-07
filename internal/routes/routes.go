package routes

import (
	"github.com/darkphotonKN/age-of-carnath/internal/server"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *server.MultiplayerServer) *gin.Engine {
	r := gin.Default()

	r.GET("/ws", server.HandleMatchSearch)

	return r
}
