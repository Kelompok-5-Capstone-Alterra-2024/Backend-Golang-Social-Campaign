package entities

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255)"`
	Token    string `gorm:"-"`
}
