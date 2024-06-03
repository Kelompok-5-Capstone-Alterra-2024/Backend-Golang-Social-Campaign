package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type TestimoniVolunteerRepository interface {
	Create(testimoniVolunteer entities.TestimoniVolunteer) (entities.TestimoniVolunteer, error)
	FindByID(id uint) (entities.TestimoniVolunteer, error)
	FindAll(page, limit int) ([]entities.TestimoniVolunteer, int, error)
	Delete(id uint) error
}

type testimoniVolunteerRepository struct {
	db *gorm.DB
}

func NewTestimoniVolunteerRepository(db *gorm.DB) TestimoniVolunteerRepository {
	return &testimoniVolunteerRepository{db: db}
}

func (r *testimoniVolunteerRepository) Create(testimoniVolunteer entities.TestimoniVolunteer) (entities.TestimoniVolunteer, error) {
	err := r.db.Create(&testimoniVolunteer).Error
	return testimoniVolunteer, err
}

func (r *testimoniVolunteerRepository) FindByID(id uint) (entities.TestimoniVolunteer, error) {
	var testimoniVolunteer entities.TestimoniVolunteer
	err := r.db.First(&testimoniVolunteer, id).Error
	return testimoniVolunteer, err
}

func (r *testimoniVolunteerRepository) FindAll(page, limit int) ([]entities.TestimoniVolunteer, int, error) {
	var testimoniVolunteers []entities.TestimoniVolunteer
	var total int64
	err := r.db.Offset((page - 1) * limit).Limit(limit).Find(&testimoniVolunteers).Count(&total).Error
	return testimoniVolunteers, int(total), err
}

func (r *testimoniVolunteerRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.TestimoniVolunteer{}, id).Error
	return err
}
