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
	FindByVacancyID(vacancyID uint, page int, limit int) ([]entities.Application, int, error)
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
	err := r.db.Preload("User").Preload("Volunteer").Where("user_id = ? AND vacancy_id = ?", customerID, vacancyID).First(&application).Error
	return application, err
}

func (r *applicationRepository) FindAll(offset, limit int) ([]entities.Application, int64, error) {
	var applications []entities.Application
	var total int64

	err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&applications).Count(&total).Error
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

func (r *applicationRepository) FindByVacancyID(vacancyID uint, page int, limit int) ([]entities.Application, int, error) {
	var applications []entities.Application
	var total int64
	offest := (page - 1) * limit

	err := r.db.Preload("User").Preload("Volunteer").Offset(offest).Limit(limit).Where("vacancy_id = ?", vacancyID).Find(&applications).Count(&total).Error
	return applications, int(total), err
}
