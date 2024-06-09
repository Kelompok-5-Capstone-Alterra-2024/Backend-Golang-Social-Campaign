package service

import (
	"capstone/entities"
	"capstone/repositories"
	"errors"

	"gorm.io/gorm"
)

type LikesCommentService interface {
	CreateLikesComment(commentID uint, userID uint) (entities.LikesComment, error)
	DeleteLikesComment(id uint) error
	GetLikesCommentByID(id uint) (entities.LikesComment, error)
	GetAllLikesComments() ([]entities.LikesComment, error)
}

type likesCommentService struct {
	repo repositories.LikesCommentRepository
}

func NewLikesCommentService(repo repositories.LikesCommentRepository) LikesCommentService {
	return &likesCommentService{repo: repo}
}

func (s *likesCommentService) CreateLikesComment(commentId, userId uint) (entities.LikesComment, error) {
	// Check if the customer already liked the comment
	existingLike, err := s.repo.FindByCustomerAndComment(userId, commentId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.LikesComment{}, err
	}
	if existingLike.ID != 0 {
		return entities.LikesComment{}, errors.New("customer already liked this comment")
	}

	like := entities.LikesComment{
		UserID:    userId,
		CommentID: commentId,
	}

	// Create the new like
	_, err = s.repo.Create(like)
	if err != nil {
		return entities.LikesComment{}, err
	}

	// Increment the like count for the comment
	err = s.repo.IncrementLike(commentId)
	if err != nil {
		return entities.LikesComment{}, err
	}

	return like, nil

}

func (s *likesCommentService) DeleteLikesComment(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	// Decrement the like count for the comment
	err = s.repo.DecrementLike(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *likesCommentService) GetLikesCommentByID(id uint) (entities.LikesComment, error) {
	return s.repo.FindByID(id)
}

func (s *likesCommentService) GetAllLikesComments() ([]entities.LikesComment, error) {
	return s.repo.FindAll()
}
