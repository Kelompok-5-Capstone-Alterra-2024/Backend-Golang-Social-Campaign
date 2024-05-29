package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Save(application entities.Application) (entities.Application, error)
}

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *applicationRepository {
	return &applicationRepository{db}
}

func (r *applicationRepository) Save(application entities.Application) (entities.Application, error) {
	if err := r.db.Create(&application).Error; err != nil {
		return application, err
	}
	return application, nil
}
