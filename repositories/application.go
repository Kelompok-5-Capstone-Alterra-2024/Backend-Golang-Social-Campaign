package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Save(application entities.Application) (entities.Application, error)
	FindByCustomerIDAndVacancyID(customerID, vacancyID uint) (entities.Application, error)
	FindAll(offset, limit int) ([]entities.Application, int64, error)
	FindByID(id uint) (entities.Application, error)
	DeleteByID(id uint) error
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

func (r *applicationRepository) FindByCustomerIDAndVacancyID(customerID, vacancyID uint) (entities.Application, error) {
	var application entities.Application
	err := r.db.Where("customer_id = ? AND vacancy_id = ?", customerID, vacancyID).First(&application).Error
	return application, err
}

func (r *applicationRepository) FindAll(offset, limit int) ([]entities.Application, int64, error) {
	var applications []entities.Application
	var total int64

	err := r.db.Offset(offset).Limit(limit).Find(&applications).Count(&total).Error
	return applications, total, err
}

func (r *applicationRepository) FindByID(id uint) (entities.Application, error) {
	var application entities.Application
	err := r.db.First(&application, id).Error
	return application, err
}

func (r *applicationRepository) DeleteByID(id uint) error {
	return r.db.Delete(&entities.Application{}, id).Error
}
