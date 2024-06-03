package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entities.Comment) (entities.Comment, error)
	Update(comment entities.Comment) (entities.Comment, error)
	FindByID(id uint) (entities.Comment, error)
	FindAll(page, limit int) ([]entities.Comment, int, error)
	Delete(id uint) error
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

func (r *commentRepository) FindAll(page, limit int) ([]entities.Comment, int, error) {
	var comments []entities.Comment
	var total int64
	err := r.db.Offset((page - 1) * limit).Limit(limit).Find(&comments).Count(&total).Error
	return comments, int(total), err
}

func (r *commentRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.Comment{}, id).Error
	return err
}
