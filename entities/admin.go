package entities

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255)"`
	Token    string `gorm:"-"`
}

type AdminNotification struct {
	gorm.Model
	UserID    uint   `json:"-"`
	UserName  string `json:"user_name" gorm:"type:varchar(255)"`
	AvatarURL string `json:"avatar_url" gorm:"type:varchar(255)"`
	Message   string `json:"message" gorm:"type:text"`
}
