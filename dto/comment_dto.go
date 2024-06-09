package dto

import (
	"capstone/entities"
)

type CommentRequest struct {
	Body string `json:"body"` // Rename to Content to match input JSON
}

func (r *CommentRequest) ToEntity(articleId, userId uint) entities.Comment {
	return entities.Comment{
		UserID:    userId,
		ArticleID: articleId,
		Comment:   r.Body, // Map Content to Comment
	}
}

type CommentResponses struct {
	ID         uint                       `json:"id"`
	User       UserCommentArticleResponse `json:"user"`
	CreatedAt  string                     `json:"created_at"`
	Body       string                     `json:"body"`
	TotalLikes int                        `json:"total_likes"`
}

type UserCommentArticleResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func ToUserCommentArticleResponse(user entities.User) UserCommentArticleResponse {
	return UserCommentArticleResponse{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}

func ToCommentResponses(comment entities.Comment) CommentResponses {
	return CommentResponses{
		ID:         comment.ID,
		User:       ToUserCommentArticleResponse(comment.User),
		CreatedAt:  comment.CreatedAt.Format("2006-01-02"),
		Body:       comment.Comment,
		TotalLikes: comment.TotalLikes,
	}
}

func ToCommentResponsesList(comments []entities.Comment) []CommentResponses {
	var res []CommentResponses
	for _, comment := range comments {
		res = append(res, ToCommentResponses(comment))
	}
	return res
}
