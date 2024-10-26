package routes

import (
	config "github.com/darkphotonKN/age-of-carnath/config/db"
	"github.com/darkphotonKN/age-of-carnath/internal/server"
	"github.com/darkphotonKN/age-of-carnath/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *server.MultiplayerServer) *gin.Engine {
	r := gin.Default()

	r.GET("/ws", server.HandleMatchConn)

	// Base Routes Prefix
	api := r.Group("/api")

	// -- Users --

	// --- User Setup ---
	userRepo := user.NewUserRepository(config.DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// --- User Routes ---
	userRoutes := api.Group("/user")
	userRoutes.POST("/signup", userHandler.SignUp)

	return r
}
