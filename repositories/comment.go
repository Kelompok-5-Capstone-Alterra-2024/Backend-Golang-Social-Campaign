package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entities.Comment) (entities.Comment, error)
	Update(comment entities.Comment) (entities.Comment, error)
	FindByID(id uint) (entities.Comment, error)
	FindAllByArticleID(id uint, page, limit int) ([]entities.Comment, int, error)
	Delete(id uint) error
	FindAll() ([]entities.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment entities.Comment) (entities.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *commentRepository) Update(comment entities.Comment) (entities.Comment, error) {
	err := r.db.Save(&comment).Error
	return comment, err
}

func (r *commentRepository) FindByID(id uint) (entities.Comment, error) {
	var comment entities.Comment
	err := r.db.First(&comment, id).Error
	return comment, err
}

func (r *commentRepository) FindAllByArticleID(id uint, page, limit int) ([]entities.Comment, int, error) {
	var comments []entities.Comment
	var total int64
	err := r.db.Model(&comments).Where("article_id = ?", id).Count(&total).Error
	if err != nil {
		return comments, 0, err
	}
	err = r.db.Limit(limit).Offset((page-1)*limit).Preload("User").Where("article_id = ?", id).Find(&comments).Error
	return comments, int(total), err
}

func (r *commentRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.Comment{}, id).Error
	return err
}

func (r *commentRepository) FindAll() ([]entities.Comment, error) {
	var comments []entities.Comment
	err := r.db.Preload("User").Find(&comments).Error
	return comments, err
}
