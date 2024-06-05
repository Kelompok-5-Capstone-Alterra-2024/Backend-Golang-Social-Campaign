package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (entities.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByUsername(username string) (entities.Admin, error) {
	var admin entities.Admin
	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}
