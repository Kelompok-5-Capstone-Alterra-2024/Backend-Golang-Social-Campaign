package repositories

import (
	"capstone/dto"

	"gorm.io/gorm"
)

type ChatbotRepository interface {
	Create(chatbot dto.Chatbot) (dto.Chatbot, error)
	FindByID(id string) ([]dto.Chatbot, error)
}

type chatbotRepository struct {
	db *gorm.DB
}

func NewChatbotRepository(db *gorm.DB) *chatbotRepository {
	return &chatbotRepository{db}
}

func (r *chatbotRepository) Create(chatbot dto.Chatbot) (dto.Chatbot, error) {
	err := r.db.Create(&chatbot).Error
	return chatbot, err
}

func (r *chatbotRepository) FindByID(id string) ([]dto.Chatbot, error) {
	var chatbot []dto.Chatbot
	err := r.db.Where("chat_id = ?", id).Order("created_at desc").Find(&chatbot).Error
	return chatbot, err
}
