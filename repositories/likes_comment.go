package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type LikesCommentRepository interface {
	Create(likesComment entities.LikesComment) (entities.LikesComment, error)
	Delete(id uint) error
	FindByCustomerAndComment(customerID, commentID uint) (entities.LikesComment, error)
	FindByID(id uint) (entities.LikesComment, error)
	FindAll() ([]entities.LikesComment, error)
}

type likesCommentRepository struct {
	db *gorm.DB
}

func NewLikesCommentRepository(db *gorm.DB) LikesCommentRepository {
	return &likesCommentRepository{db: db}
}

func (r *likesCommentRepository) FindByCustomerAndComment(customerID, commentID uint) (entities.LikesComment, error) {
	var like entities.LikesComment
	if err := r.db.Where("customer_id = ? AND comment_id = ?", customerID, commentID).First(&like).Error; err != nil {
		return entities.LikesComment{}, err
	}
	return like, nil
}

func (r *likesCommentRepository) Create(likesComment entities.LikesComment) (entities.LikesComment, error) {
	err := r.db.Create(&likesComment).Error
	return likesComment, err
}

func (r *likesCommentRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.LikesComment{}, id).Error
	return err
}

func (r *likesCommentRepository) FindByID(id uint) (entities.LikesComment, error) {
	var like entities.LikesComment
	if err := r.db.First(&like, id).Error; err != nil {
		return entities.LikesComment{}, err
	}
	return like, nil
}

func (r *likesCommentRepository) FindAll() ([]entities.LikesComment, error) {
	var likes []entities.LikesComment
	if err := r.db.Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
