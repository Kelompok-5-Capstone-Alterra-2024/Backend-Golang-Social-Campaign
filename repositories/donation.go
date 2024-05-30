package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type DonationRepository interface {
	Save(donation entities.Donation) (entities.Donation, error)
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
