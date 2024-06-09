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
		UserID:    r.CustomerID,
		CommentID: r.CommentID,
	}
}

type LikesCommentResponse struct {
	ID         uint `json:"id"`
	CustomerID uint `json:"customer_id"`
	CommentID  uint `json:"comment_id"`
}
