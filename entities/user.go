package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname   string `json:"fullname" gorm:"type:varchar(255)"`
	Username   string `json:"username" gorm:"type:varchar(255)"`
	Email      string `json:"email" gorm:"type:varchar(255)"`
	Password   string `json:"password" gorm:"type:varchar(255)"`
	Avatar     string `json:"avatar" gorm:"type:varchar(255)"`
	Token      string `gorm:"-"`
	ResetToken string `gorm:"type:varchar(255)"`
}
