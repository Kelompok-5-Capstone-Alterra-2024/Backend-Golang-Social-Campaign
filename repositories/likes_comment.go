package repositories

import (
	"capstone/entities"
	"context"

	"gorm.io/gorm"
)

type LikesCommentRepository interface {
	Create(ctx context.Context, likesComment entities.LikesComment) error
	Delete(ctx context.Context, commentID uint, userID uint) error
	FindByCustomerAndComment(customerID, commentID uint) (entities.LikesComment, error)
	FindByID(id uint) (entities.LikesComment, error)
	FindAll() ([]entities.LikesComment, error)
	IsLiked(commentID uint, userID uint) (bool, error)
	IncrementLike(ctx context.Context, commentID uint) error
	DecrementLike(ctx context.Context, commentID uint) error
}

type likesCommentRepository struct {
	db *gorm.DB
}

func NewLikesCommentRepository(db *gorm.DB) LikesCommentRepository {
	return &likesCommentRepository{db: db}
}

func (r *likesCommentRepository) FindByCustomerAndComment(customerID, commentID uint) (entities.LikesComment, error) {
	var like entities.LikesComment
	if err := r.db.Where("user_id = ? AND comment_id = ?", customerID, commentID).First(&like).Error; err != nil {
		return entities.LikesComment{}, err
	}
	return like, nil
}

func (r *likesCommentRepository) Create(ctx context.Context, likesComment entities.LikesComment) error {
	if err := r.db.Create(&likesComment).Error; err != nil {
		return err
	}
	return nil
}

func (r *likesCommentRepository) Delete(ctx context.Context, commentID uint, userID uint) error {
	return r.db.WithContext(ctx).Where("comment_id = ? AND user_id = ?", commentID, userID).Delete(&entities.LikesComment{}).Error
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

func (r *likesCommentRepository) IsLiked(commentID uint, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entities.LikesComment{}).Where("comment_id = ? AND user_id = ?", commentID, userID).Count(&count).Error
	return count > 0, err
}

func (r *likesCommentRepository) IncrementLike(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Model(&entities.Comment{}).Where("id = ?", commentID).UpdateColumn("total_likes", gorm.Expr("total_likes + ?", 1)).Error
}

func (r *likesCommentRepository) DecrementLike(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Model(&entities.Comment{}).Where("id = ?", commentID).UpdateColumn("total_likes", gorm.Expr("total_likes - ?", 1)).Error
}
