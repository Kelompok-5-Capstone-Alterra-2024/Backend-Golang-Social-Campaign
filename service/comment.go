package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type CommentService interface {
	CreateComment(comment entities.Comment) (entities.Comment, error)
	UpdateComment(comment entities.Comment) (entities.Comment, error)
	FindByID(id uint) (entities.Comment, error)
	GetAllByArticleID(id uint, page, limit int) ([]entities.Comment, int, error)
	DeleteComment(id uint) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) CreateComment(comment entities.Comment) (entities.Comment, error) {
	return s.repo.Create(comment)
}

func (s *commentService) UpdateComment(comment entities.Comment) (entities.Comment, error) {
	return s.repo.Update(comment)
}

func (s *commentService) FindByID(id uint) (entities.Comment, error) {
	return s.repo.FindByID(id)
}

func (s *commentService) GetAllByArticleID(id uint, page, limit int) ([]entities.Comment, int, error) {
	return s.repo.FindAllByArticleID(id, page, limit)
}

func (s *commentService) DeleteComment(id uint) error {
	return s.repo.Delete(id)
}
