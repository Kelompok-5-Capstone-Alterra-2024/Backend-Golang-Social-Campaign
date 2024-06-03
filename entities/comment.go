package entities

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CustomerID uint   `json:"customer_id" gorm:"type:bigint"`
	ArticleID  uint   `json:"article_id" gorm:"type:bigint"`
	Comment    string `json:"comment" gorm:"type:varchar(255)"`
}
