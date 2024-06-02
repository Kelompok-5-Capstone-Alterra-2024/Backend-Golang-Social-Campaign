package entities

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	AdminID  uint   `json:"admin_id" gorm:"type:bigint"`
	Title    string `json:"title" gorm:"type:varchar(255)"`
	Content  string `json:"content" gorm:"type:text"`
	ImageURL string `json:"image_url" gorm:"type:varchar(255)"`
}
