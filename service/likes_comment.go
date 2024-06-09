package service

import (
	"capstone/entities"
	"capstone/repositories"
	"errors"

	"gorm.io/gorm"
)

type LikesCommentService interface {
	CreateLikesComment(likesComment entities.LikesComment) (entities.LikesComment, error)
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

func (s *likesCommentService) CreateLikesComment(likesComment entities.LikesComment) (entities.LikesComment, error) {
	// Check if the customer already liked the comment
	existingLike, err := s.repo.FindByCustomerAndComment(likesComment.UserID, likesComment.CommentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.LikesComment{}, err
	}
	if existingLike.ID != 0 {
		return entities.LikesComment{}, errors.New("customer already liked this comment")
	}

	return s.repo.Create(likesComment)
}

func (s *likesCommentService) DeleteLikesComment(id uint) error {
	return s.repo.Delete(id)
}

func (s *likesCommentService) GetLikesCommentByID(id uint) (entities.LikesComment, error) {
	return s.repo.FindByID(id)
}

func (s *likesCommentService) GetAllLikesComments() ([]entities.LikesComment, error) {
	return s.repo.FindAll()
}
