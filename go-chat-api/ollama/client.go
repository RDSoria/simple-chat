package ollama

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/go-resty/resty/v2"
)

type OllamaResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) SendMessage(message string, senderLang string, receiverLang string) (string, error) {
	client := resty.New()
	url := "http://localhost:11434/api/chat"

	payload := map[string]interface{}{
		"model": "deepseek-r1:1.5b", //"gemma2:2b" ,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "could you translate this message written in " + senderLang + " to " + receiverLang + "? the message '" + message + "' **important** only response with the translated message",
			},
		},
		"stream": false,
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(url)

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New("error from Ollama API: " + resp.Status())
	}

	// Log the raw response for debugging
	log.Printf("Raw response from Ollama: %s", resp.String())

	// Parse the JSON response
	var ollamaResponse OllamaResponse
	if err := json.Unmarshal(resp.Body(), &ollamaResponse); err != nil {
		return "", errors.New("failed to parse response from Ollama")
	}

	return ollamaResponse.Message.Content, nil
}
