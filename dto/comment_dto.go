package dto

import (
	"capstone/entities"
)

type CommentRequest struct {
	CustomerID uint   `json:"user_id"` // Mapping user_id to CustomerID
	ArticleID  uint   `json:"article_id"`
	Content    string `json:"content"` // Rename to Content to match input JSON
}

func (r *CommentRequest) ToEntity() entities.Comment {
	return entities.Comment{
		CustomerID: r.CustomerID,
		ArticleID:  r.ArticleID,
		Comment:    r.Content, // Map Content to Comment
	}
}
