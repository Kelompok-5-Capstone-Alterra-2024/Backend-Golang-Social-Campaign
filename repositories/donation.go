package repositories

import (
	"capstone/entities"
	"context"

	"gorm.io/gorm"
)

type DonationRepository interface {
	Save(donation entities.Donation) (entities.Donation, error)
	SaveComment(donationComment entities.DonationComment) (entities.DonationComment, error)
	GetCommentsByFundraisingID(id int) ([]entities.DonationComment, error)
	GetByUserID(limit int, offset int, userID uint) ([]entities.Donation, error)
	GetByID(id int) (entities.Donation, error)
	GetByFundraisingID(id int) ([]entities.Donation, error)
	GetByCode(code string) (entities.Donation, error)
	Update(donation entities.Donation) (entities.Donation, error)
	CreateLike(ctx context.Context, like entities.LikeDonationComment) error
	RemoveLike(ctx context.Context, commentID uint, userID uint) error
	IsLiked(ctx context.Context, commentID uint, userID uint) (bool, error)
	IncrementLike(ctx context.Context, commentID uint) error
	DecrementLike(ctx context.Context, commentID uint) error
}

type donationRepository struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) *donationRepository {
	return &donationRepository{db}
}

func (r *donationRepository) Save(donation entities.Donation) (entities.Donation, error) {
	if err := r.db.Create(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationRepository) SaveComment(donationComment entities.DonationComment) (entities.DonationComment, error) {
	if err := r.db.Create(&donationComment).Error; err != nil {
		return donationComment, err
	}
	return donationComment, nil
}

func (r *donationRepository) GetCommentsByFundraisingID(id int) ([]entities.DonationComment, error) {
	var donationComments []entities.DonationComment
	if err := r.db.Preload("Donation.User").Preload("Donation.Fundraising").Joins("JOIN donations ON donations.id = donation_comments.donation_id").Where("donations.fundraising_id = ?", id).Find(&donationComments).Error; err != nil {
		return donationComments, err
	}
	return donationComments, nil
}

func (r *donationRepository) GetByUserID(limit int, offset int, userID uint) ([]entities.Donation, error) {
	var donations []entities.Donation
	if err := r.db.Preload("Fundraising").Where("user_id = ?", userID).Order("created_at desc").Limit(limit).Offset(offset).Find(&donations).Error; err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *donationRepository) GetByID(id int) (entities.Donation, error) {
	var donation entities.Donation
	if err := r.db.Preload("Fundraising.Organization").Where("id = ?", id).First(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationRepository) GetByFundraisingID(id int) ([]entities.Donation, error) {
	var donations []entities.Donation
	if err := r.db.Preload("User").Where("fundraising_id = ?", id).Find(&donations).Error; err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *donationRepository) GetByCode(code string) (entities.Donation, error) {
	var donation entities.Donation
	if err := r.db.Where("code = ?", code).First(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationRepository) Update(donation entities.Donation) (entities.Donation, error) {
	if err := r.db.Save(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *donationRepository) CreateLike(ctx context.Context, like entities.LikeDonationComment) error {
	if err := r.db.Create(&like).Error; err != nil {
		return err
	}
	return nil
}

func (r *donationRepository) RemoveLike(ctx context.Context, commentID uint, userID uint) error {
	if err := r.db.Where("donation_comment_id = ? AND user_id = ?", commentID, userID).Delete(&entities.LikeDonationComment{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *donationRepository) IsLiked(ctx context.Context, commentID uint, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entities.LikeDonationComment{}).Where("donation_comment_id = ? AND user_id = ?", commentID, userID).Count(&count).Error
	return count > 0, err
}

func (r *donationRepository) IncrementLike(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Model(&entities.DonationComment{}).Where("id = ?", commentID).UpdateColumn("total_likes", gorm.Expr("total_likes + ?", 1)).Error
}

func (r *donationRepository) DecrementLike(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Model(&entities.DonationComment{}).Where("id = ?", commentID).UpdateColumn("total_likes", gorm.Expr("total_likes - ?", 1)).Error
}
