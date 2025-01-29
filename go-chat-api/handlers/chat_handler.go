package handlers

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/RDSoria/simple-chat/ollama"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	OllamaClient *ollama.Client
}

func NewChatHandler(client *ollama.Client) *ChatHandler {
	return &ChatHandler{OllamaClient: client}
}

// Message structure
type Message struct {
	OriginalText   string `json:"original_text"`
	TranslatedText string `json:"translated_text"`
	Lang           string `json:"lang"`
	User           string `json:"user"`
}

var (
	messages      []Message
	messagesMutex sync.Mutex
)

// SendMessageHandler handles incoming messages and appends a translation note
func (h *ChatHandler) SendMessageHandler(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Determine languages based on the sender
	senderLang, receiverLang := GetTranslationLanguages(message.User)

	//translate
	// Send the message to the Ollama API
	response, err := h.OllamaClient.SendMessage(message.OriginalText, senderLang, receiverLang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	// if you want to avoid deepseek thinking
	start := strings.Index(response, "</think>")
	if start != -1 {
		response = response[start+len("</think>"):]
	}

	// Trim any leading or trailing whitespace
	message.TranslatedText = " [" + senderLang + " > " + receiverLang + "] : " + strings.TrimSpace(response)

	//translate

	// Store message
	messagesMutex.Lock()
	messages = append(messages, message)
	messagesMutex.Unlock()

	log.Printf("Message received: %+v\n", message)
	c.JSON(http.StatusOK, gin.H{"status": "Message sent"})
}

// GetMessagesHandler returns all stored messages
func GetMessagesHandler(c *gin.Context) {
	messagesMutex.Lock()
	defer messagesMutex.Unlock()

	c.JSON(http.StatusOK, messages)
}
