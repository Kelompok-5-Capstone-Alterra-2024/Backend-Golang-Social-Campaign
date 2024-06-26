package dto

import "time"

type Chatbot struct {
	ID        string
	ChatID    string
	Role      string
	Message   string
	CreatedAt time.Time
}

type Message struct {
	Role    string
	Message string
}

type ChatBotRequest struct {
	ID      string
	Message string
}

type ChatbotResponse struct {
	ID        string `json:"id"`
	ChatID    string `json:"chat_id"`
	Role      string `json:"role"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}
