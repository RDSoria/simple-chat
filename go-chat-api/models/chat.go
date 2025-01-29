package models

// type MessageRequest struct {
// 	Role    string `json:"role"`
// 	Content string `json:"content"`
// }

//	type ChatRequest struct {
//		Messages []MessageRequest `json:"messages" binding:"required"`
//		Stream   bool             `json:"stream"`
//	}
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}
