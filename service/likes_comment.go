package service

import (
	"capstone/entities"
	"capstone/repositories"
	"context"
)

type LikesCommentService interface {
	LikesComment(ctx context.Context, commentID uint, userID uint) error
	DeleteLikesComment(ctx context.Context, commentID uint, userID uint) error
	GetLikesCommentByID(id uint) (entities.LikesComment, error)
	GetAllLikesComments() ([]entities.LikesComment, error)
}

type likesCommentService struct {
	repo repositories.LikesCommentRepository
}

func NewLikesCommentService(repo repositories.LikesCommentRepository) LikesCommentService {
	return &likesCommentService{repo: repo}
}

func (s *likesCommentService) LikesComment(ctx context.Context, commentId, userId uint) error {
	// Check if the customer already liked the comment
	liked, err := s.repo.IsLiked(commentId, userId)
	if err != nil {
		return err
	}
	if liked {
		return nil
	}

	like := entities.LikesComment{
		CommentID: commentId,
		UserID:    userId,
	}

	err = s.repo.Create(ctx, like)
	if err != nil {
		return err
	}

	return s.repo.IncrementLike(ctx, commentId)

}

func (s *likesCommentService) DeleteLikesComment(ctx context.Context, commentID uint, userID uint) error {
	err := s.repo.Delete(ctx, commentID, userID)
	if err != nil {
		return err
	}

	return s.repo.DecrementLike(ctx, commentID)

}

func (s *likesCommentService) GetLikesCommentByID(id uint) (entities.LikesComment, error) {
	return s.repo.FindByID(id)
}

func (s *likesCommentService) GetAllLikesComments() ([]entities.LikesComment, error) {
	return s.repo.FindAll()
}
