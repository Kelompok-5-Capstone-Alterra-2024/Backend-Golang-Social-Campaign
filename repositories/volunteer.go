package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type VolunteerRepository interface {
	FindAll() ([]entities.Volunteer, error)
	FindByID(id uint) (entities.Volunteer, error)
	Save(volunteer entities.Volunteer) (entities.Volunteer, error)
}

type volunteerRepository struct {
	db *gorm.DB
}

func NewVolunteerRepository(db *gorm.DB) *volunteerRepository {
	return &volunteerRepository{db}
}

func (r *volunteerRepository) FindAll() ([]entities.Volunteer, error) {
	var volunteers []entities.Volunteer
	if err := r.db.Find(&volunteers).Error; err != nil {
		return volunteers, err
	}
	return volunteers, nil
}

func (r *volunteerRepository) FindByID(id uint) (entities.Volunteer, error) {
	var volunteer entities.Volunteer
	if err := r.db.Where("id = ?", id).First(&volunteer).Error; err != nil {
		return volunteer, err
	}
	return volunteer, nil
}

func (r *volunteerRepository) Save(volunteer entities.Volunteer) (entities.Volunteer, error) {
	if err := r.db.Create(&volunteer).Error; err != nil {
		return volunteer, err
	}
	return volunteer, nil
}
