package entities

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string `json:"title" gorm:"type:varchar(255)"`
	Content  string `json:"content" gorm:"type:text"`
	ImageURL string `json:"image_url" gorm:"type:varchar(255)"`
}
