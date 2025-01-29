package routes

import (
	"github.com/RDSoria/simple-chat/handlers"
	"github.com/RDSoria/simple-chat/ollama"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API endpoints
func RegisterRoutes(r *gin.Engine) {
	ollamaClient := ollama.NewClient("")
	chatHandler := handlers.NewChatHandler(ollamaClient)

	r.POST("/api/messages", chatHandler.SendMessageHandler)
	r.GET("/api/messages", handlers.GetMessagesHandler)
	r.POST("/api/set-language", handlers.SetLanguageHandler)
}
