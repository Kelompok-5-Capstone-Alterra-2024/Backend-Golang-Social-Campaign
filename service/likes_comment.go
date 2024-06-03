package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type LikesCommentService interface {
	CreateLikesComment(likesComment entities.LikesComment) (entities.LikesComment, error)
	DeleteLikesComment(id uint) error
}

type likesCommentService struct {
	repo repositories.LikesCommentRepository
}

func NewLikesCommentService(repo repositories.LikesCommentRepository) LikesCommentService {
	return &likesCommentService{repo: repo}
}

func (s *likesCommentService) CreateLikesComment(likesComment entities.LikesComment) (entities.LikesComment, error) {
	return s.repo.Create(likesComment)
}

func (s *likesCommentService) DeleteLikesComment(id uint) error {
	return s.repo.Delete(id)
}
