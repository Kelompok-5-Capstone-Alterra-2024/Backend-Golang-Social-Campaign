package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type VolunteerRepository interface {
	FindAll(page, limit int) ([]entities.Volunteer, int64, error)
	FindByID(id uint) (entities.Volunteer, error)
	Save(volunteer entities.Volunteer) (entities.Volunteer, error)
}

type volunteerRepository struct {
	db *gorm.DB
}

func NewVolunteerRepository(db *gorm.DB) *volunteerRepository {
	return &volunteerRepository{db}
}

func (r *volunteerRepository) FindAll(page, limit int) ([]entities.Volunteer, int64, error) {
	var volunteers []entities.Volunteer
	var totalCount int64

	offset := (page - 1) * limit

	if err := r.db.Offset(offset).Limit(limit).Find(&volunteers).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	return volunteers, totalCount, nil
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
