package dto

import (
	"capstone/entities"
)

type LikesCommentRequest struct {
	CustomerID uint `json:"customer_id"`
	CommentID  uint `json:"comment_id"`
}

func (r *LikesCommentRequest) ToEntity() entities.LikesComment {
	return entities.LikesComment{
		CustomerID: r.CustomerID,
		CommentID:  r.CommentID,
	}
}
