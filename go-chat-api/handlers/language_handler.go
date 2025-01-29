package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userLanguages  = map[string]string{"A": "English", "B": "Spanish"}
	languagesMutex sync.Mutex
)

// SetLanguageHandler updates the language for a user
func SetLanguageHandler(c *gin.Context) {
	var data struct {
		User string `json:"user"`
		Lang string `json:"lang"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	languagesMutex.Lock()
	userLanguages[data.User] = data.Lang
	languagesMutex.Unlock()

	log.Printf("Updated language for user %s to %s\n", data.User, data.Lang)
	c.JSON(http.StatusOK, gin.H{"status": "Language updated"})
}

// GetTranslationLanguages retrieves the sender and receiver languages
func GetTranslationLanguages(user string) (senderLang, receiverLang string) {
	languagesMutex.Lock()
	defer languagesMutex.Unlock()

	senderLang = userLanguages[user]
	receiverLang = userLanguages["A"]
	if user == "A" {
		receiverLang = userLanguages["B"]
	}
	return
}
