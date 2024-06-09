package entities

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID    uint    `json:"-"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	ArticleID uint    `json:"-"`
	Article   Article `json:"article" gorm:"foreignKey:ArticleID"`
	Comment   string  `json:"comment" gorm:"type:varchar(255)"`
}
