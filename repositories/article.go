package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article entities.Article) (entities.Article, error)
	Update(article entities.Article) (entities.Article, error)
	FindByID(id uint) (entities.Article, error)
	FindAll(page, limit int) ([]entities.Article, int, error)
	Delete(id uint) error
	Save(article *entities.Article) *gorm.DB
	FindTop() ([]entities.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(article entities.Article) (entities.Article, error) {
	err := r.db.Create(&article).Error
	return article, err
}

func (r *articleRepository) Update(article entities.Article) (entities.Article, error) {
	err := r.db.Save(&article).Error
	return article, err
}

func (r *articleRepository) FindByID(id uint) (entities.Article, error) {
	var article entities.Article
	err := r.db.First(&article, id).Error
	return article, err
}

func (r *articleRepository) FindAll(page, limit int) ([]entities.Article, int, error) {
	var articles []entities.Article
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	r.db.Model(&entities.Article{}).Count(&total)

	return articles, int(total), nil
}

func (r *articleRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Article{}, id).Error
}

func (r *articleRepository) Save(article *entities.Article) *gorm.DB {
	return r.db.Save(article)
}

func (r *articleRepository) FindTop() ([]entities.Article, error) {
	var articles []entities.Article
	err := r.db.Table("articles").
		Select("articles.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.article_id = articles.id").
		Group("articles.id").
		Order("comment_count DESC").
		Limit(2).
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}
