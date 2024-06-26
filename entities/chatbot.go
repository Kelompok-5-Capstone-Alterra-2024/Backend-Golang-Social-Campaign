package entities

import "gorm.io/gorm"

type Chatbot struct {
	gorm.Model
	ID      string `json:"id" gorm:"primaryKey"`
	ChatID  string `json:"chat_id" gorm:"type:varchar(255)"`
	Role    string `json:"role" gorm:"type:varchar(255)"`
	Message string `json:"message" gorm:"type:text"`
}
