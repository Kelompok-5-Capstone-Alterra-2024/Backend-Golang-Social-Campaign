package entities

import (
	"gorm.io/gorm"
)

type LikesComment struct {
	gorm.Model
	CustomerID uint `json:"customer_id" gorm:"type:bigint"`
	CommentID  uint `json:"comment_id" gorm:"type:bigint"`
}
