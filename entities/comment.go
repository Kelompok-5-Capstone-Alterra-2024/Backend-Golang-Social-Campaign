package entities

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID     uint    `json:"-"`
	User       User    `json:"user" gorm:"foreignKey:UserID"`
	ArticleID  uint    `json:"-"`
	Article    Article `json:"article" gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comment    string  `json:"comment" gorm:"type:varchar(255)"`
	TotalLikes int     `json:"total_likes" gorm:"type:int"`
}
