package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type ArticleService interface {
	CreateArticle(article entities.Article) (entities.Article, error)
	UpdateArticle(article entities.Article) (entities.Article, error)
	FindByID(id uint) (entities.Article, error)
	FindAll(page, limit int) ([]entities.Article, int, error)
	DeleteArticle(id uint) error
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) CreateArticle(article entities.Article) (entities.Article, error) {
	return s.repo.Create(article)
}

func (s *articleService) UpdateArticle(article entities.Article) (entities.Article, error) {
	existingArticle, err := s.repo.FindByID(article.ID)
	if err != nil {
		return entities.Article{}, err
	}

	existingArticle.Title = article.Title
	existingArticle.Content = article.Content
	existingArticle.ImageURL = article.ImageURL
	existingArticle.AdminID = article.AdminID

	if err := s.repo.Save(&existingArticle).Error; err != nil {
		return entities.Article{}, err
	}

	return existingArticle, nil
}

func (s *articleService) FindByID(id uint) (entities.Article, error) {
	return s.repo.FindByID(id)
}

func (s *articleService) FindAll(page, limit int) ([]entities.Article, int, error) {
	return s.repo.FindAll(page, limit)
}

func (s *articleService) DeleteArticle(id uint) error {
	return s.repo.Delete(id)
}
