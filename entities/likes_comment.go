package entities

import (
	"gorm.io/gorm"
)

type LikesComment struct {
	gorm.Model
	UserID    uint    `json:"-"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	CommentID uint    `json:"-"`
	Comment   Comment `json:"comment" gorm:"foreignKey:CommentID"`
}
