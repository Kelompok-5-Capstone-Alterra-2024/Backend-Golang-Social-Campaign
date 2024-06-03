package dto

import (
	"capstone/entities"
)

type CommentRequest struct {
	CustomerID uint   `json:"customer_id"`
	ArticleID  uint   `json:"article_id"`
	Comment    string `json:"comment"`
}

func (r *CommentRequest) ToEntity() entities.Comment {
	return entities.Comment{
		CustomerID: r.CustomerID,
		ArticleID:  r.ArticleID,
		Comment:    r.Comment,
	}
}
