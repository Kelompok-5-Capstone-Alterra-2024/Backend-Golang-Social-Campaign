package entities

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	Avatar      string `json:"avatar" gorm:"type:varchar(255)"`
	IsVerified  bool   `json:"is_verified"`
}
