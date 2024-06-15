package repositories

import (
	"capstone/entities"
	"context"

	"gorm.io/gorm"
)

type DonationManualRepository interface {
	Save(donation entities.DonationManual) (entities.DonationManual, error)
	SaveComment(donationComment entities.DonationManualComment) (entities.DonationManualComment, error)
	GetCommentsByDonationID(id int) ([]entities.DonationManualComment, error)
	GetByUserID(limit int, offset int, userID uint) ([]entities.DonationManual, error)
	GetByID(id int) (entities.DonationManual, error)
	GetByDonationID(id int) ([]entities.DonationManual, error)
	Update(donation entities.DonationManual) (entities.DonationManual, error)
	CreateLike(ctx context.Context, like entities.LikeDonationManualComment) error
	RemoveLike(ctx context.Context, commentID uint, userID uint) error
	IsLiked(ctx context.Context, commentID uint, userID uint) (bool, error)
	IncrementLike(ctx context.Context, commentID uint) error
	DecrementLike(ctx context.Context, commentID uint) error
}

type donationManualRepository struct {
	db *gorm.DB
}

func NewDonationManualRepository(db *gorm.DB) *donationManualRepository {
	return &donationManualRepository{db}
}

func (r *donationManualRepository) Save(donation entities.DonationManual) (entities.DonationManual, error) {
	if err := r.db.Create(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationManualRepository) SaveComment(donationComment entities.DonationManualComment) (entities.DonationManualComment, error) {
	if err := r.db.Create(&donationComment).Error; err != nil {
		return donationComment, err
	}
	return donationComment, nil
}

func (r *donationManualRepository) GetCommentsByDonationID(id int) ([]entities.DonationManualComment, error) {
	var donationComments []entities.DonationManualComment
	if err := r.db.Preload("DonationManual.User").Preload("DonationManual.Fundraising").Joins("JOIN donation_manuals ON donation_manuals.id = donation_manual_comments.donation_manual_id").Where("donation_manuals.fundraising_id = ?", id).Find(&donationComments).Error; err != nil {
		return donationComments, err
	}

	return donationComments, nil
}

func (r *donationManualRepository) GetByUserID(limit int, offset int, userID uint) ([]entities.DonationManual, error) {
	var donations []entities.DonationManual
	if err := r.db.Preload("Fundraising").Where("user_id = ?", userID).Order("created_at desc").Limit(limit).Offset(offset).Find(&donations).Error; err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *donationManualRepository) GetByID(id int) (entities.DonationManual, error) {
	var donation entities.DonationManual
	if err := r.db.Preload("Fundraising.Organization").Where("id = ?", id).Find(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationManualRepository) GetByDonationID(id int) ([]entities.DonationManual, error) {
	var donations []entities.DonationManual
	if err := r.db.Preload("User").Where("fundraising_id = ?", id).Find(&donations).Error; err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *donationManualRepository) Update(donation entities.DonationManual) (entities.DonationManual, error) {
	if err := r.db.Save(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationManualRepository) CreateLike(ctx context.Context, like entities.LikeDonationManualComment) error {
	return r.db.WithContext(ctx).Create(&like).Error
}

func (r *donationManualRepository) RemoveLike(ctx context.Context, commentID uint, userID uint) error {
	return r.db.WithContext(ctx).Where("donation_manual_comment_id = ? AND user_id = ?", commentID, userID).Delete(&entities.LikeDonationManualComment{}).Error
}

func (r *donationManualRepository) IsLiked(ctx context.Context, commentID uint, userID uint) (bool, error) {
	var like entities.LikeDonationManualComment
	if err := r.db.WithContext(ctx).Where("donation_manual_comment_id = ? AND user_id = ?", commentID, userID).First(&like).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *donationManualRepository) IncrementLike(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Model(&entities.DonationManualComment{}).Where("id = ?", commentID).UpdateColumn("total_likes", gorm.Expr("total_likes + ?", 1)).Error
}

func (r *donationManualRepository) DecrementLike(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Model(&entities.DonationManualComment{}).Where("id = ?", commentID).UpdateColumn("total_likes", gorm.Expr("total_likes - ?", 1)).Error
}
