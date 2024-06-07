package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type VolunteerRepository interface {
	Create(volunteer entities.Volunteer) (entities.Volunteer, error)
	FindByID(id uint) (entities.Volunteer, error)
	FindAll(page int, limit int) ([]entities.Volunteer, int, error)
	FindTop() ([]entities.Volunteer, error)
	Update(volunteer entities.Volunteer) (entities.Volunteer, error)
	Delete(id uint) error
	FindApplicationByVolunteerAndCustomer(volunteerID, customerID uint) (entities.Application, error)
	UpdateByID(id uint, volunteer entities.Volunteer) (entities.Volunteer, error)
}

type volunteerRepository struct {
	db *gorm.DB
}

func NewVolunteerRepository(db *gorm.DB) *volunteerRepository {
	return &volunteerRepository{db}
}

func (r *volunteerRepository) Create(volunteer entities.Volunteer) (entities.Volunteer, error) {
	if err := r.db.Create(&volunteer).Error; err != nil {
		return volunteer, err
	}
	return volunteer, nil
}

func (r *volunteerRepository) FindByID(id uint) (entities.Volunteer, error) {
	var volunteer entities.Volunteer
	err := r.db.Preload("Organization").First(&volunteer, id).Error
	return volunteer, err
}

func (r *volunteerRepository) FindAll(page int, limit int) ([]entities.Volunteer, int, error) {
	var volunteers []entities.Volunteer
	var total int64
	offset := (page - 1) * limit

	err := r.db.Preload("Organization").Offset(offset).Limit(limit).Find(&volunteers).Count(&total).Error
	return volunteers, int(total), err
}

func (r *volunteerRepository) FindTop() ([]entities.Volunteer, error) {
	var volunteers []entities.Volunteer
	err := r.db.Preload("Organization").Order("registered_volunteer desc").Limit(3).Find(&volunteers).Error
	return volunteers, err
}

func (r *volunteerRepository) Update(volunteer entities.Volunteer) (entities.Volunteer, error) {
	if err := r.db.Save(&volunteer).Error; err != nil {
		return volunteer, err
	}
	return volunteer, nil
}

func (r *volunteerRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Volunteer{}, id).Error
}

func (r *volunteerRepository) FindApplicationByVolunteerAndCustomer(volunteerID, customerID uint) (entities.Application, error) {
	var application entities.Application
	err := r.db.Where("volunteer_id = ? AND customer_id = ?", volunteerID, customerID).First(&application).Error
	return application, err
}

func (r *volunteerRepository) UpdateByID(id uint, volunteer entities.Volunteer) (entities.Volunteer, error) {
	if err := r.db.Model(&volunteer).Where("id = ?", id).Updates(&volunteer).Error; err != nil {
		return entities.Volunteer{}, err
	}
	return volunteer, nil
}
