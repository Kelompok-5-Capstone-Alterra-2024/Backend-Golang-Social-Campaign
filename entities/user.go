package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255)"`
}
