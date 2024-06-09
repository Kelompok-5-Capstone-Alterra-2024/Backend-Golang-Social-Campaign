package dto

import (
	"capstone/entities"
)

type LikesCommentRequest struct {
	UserID    uint `json:"user_id"`
	CommentID uint `json:"comment_id"`
}

func (r *LikesCommentRequest) ToEntity(commentId, userId uint) entities.LikesComment {
	return entities.LikesComment{
		UserID:    commentId,
		CommentID: userId,
	}
}

type LikesCommentResponse struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id"`
	CommentID uint `json:"comment_id"`
}
