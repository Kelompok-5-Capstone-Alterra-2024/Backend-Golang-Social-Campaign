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
	FindAllByVacancyID(volunteerID uint) ([]entities.TestimoniVolunteer, error)
	CustomerJoinedVolunteer(customerID, volunteerID uint) (bool, error)
	HasCustomerGivenTestimony(customerID, volunteerID uint) (bool, error)
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
	err := r.db.Preload("User").Preload("Volunteer").First(&testimoniVolunteer, id).Error
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

func (r *testimoniVolunteerRepository) FindAllByVacancyID(volunteerID uint) ([]entities.TestimoniVolunteer, error) {
	var testimoniVolunteers []entities.TestimoniVolunteer

	err := r.db.Preload("User").Preload("Volunteer").Order("created_at desc").Where("vacancy_id = ?", volunteerID).Find(&testimoniVolunteers).Error

	return testimoniVolunteers, err
}

func (r *testimoniVolunteerRepository) CustomerJoinedVolunteer(customerID, volunteerID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entities.Application{}).Where("user_id = ? AND vacancy_id = ?", customerID, volunteerID).Count(&count).Error
	return count > 0, err
}

func (r *testimoniVolunteerRepository) HasCustomerGivenTestimony(customerID, volunteerID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entities.TestimoniVolunteer{}).Where("user_id = ? AND vacancy_id = ?", customerID, volunteerID).Count(&count).Error
	return count > 0, err
}
