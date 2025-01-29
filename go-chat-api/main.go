package main

import (
	"log"

	"github.com/RDSoria/simple-chat/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Allow CORS for frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Register all routes
	routes.RegisterRoutes(r)

	log.Println("Server is running on :9090")
	r.Run(":9090")
}
