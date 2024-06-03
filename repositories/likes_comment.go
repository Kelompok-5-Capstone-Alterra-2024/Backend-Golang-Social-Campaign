package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type LikesCommentRepository interface {
	Create(likesComment entities.LikesComment) (entities.LikesComment, error)
	Delete(id uint) error
}

type likesCommentRepository struct {
	db *gorm.DB
}

func NewLikesCommentRepository(db *gorm.DB) LikesCommentRepository {
	return &likesCommentRepository{db: db}
}

func (r *likesCommentRepository) Create(likesComment entities.LikesComment) (entities.LikesComment, error) {
	err := r.db.Create(&likesComment).Error
	return likesComment, err
}

func (r *likesCommentRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.LikesComment{}, id).Error
	return err
}
