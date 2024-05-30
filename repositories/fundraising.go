package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type FundraisingRepository interface {
	FindAll(limit int, offset int) ([]entities.Fundraising, error)
	FindByID(id int) (entities.Fundraising, error)
	FindAllCategories() ([]entities.FundraisingCategory, error)
	FindByCategoryID(id int) ([]entities.Fundraising, error)
}

type fundraisingRepository struct {
	db *gorm.DB
}

func NewFundraisingRepository(db *gorm.DB) *fundraisingRepository {
	return &fundraisingRepository{db}
}

func (r *fundraisingRepository) FindAll(limit int, offset int) ([]entities.Fundraising, error) {
	var fundraisings []entities.Fundraising
	if err := r.db.Preload("FundraisingCategory").Preload("Organization").Limit(limit).Offset(offset).Find(&fundraisings).Error; err != nil {
		return []entities.Fundraising{}, err
	}
	return fundraisings, nil
}

func (r *fundraisingRepository) FindByID(id int) (entities.Fundraising, error) {
	var fundraising entities.Fundraising
	if err := r.db.Preload("FundraisingCategory").Preload("Organization").Where("id = ?", id).First(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}

	return fundraising, nil
}

func (r *fundraisingRepository) FindAllCategories() ([]entities.FundraisingCategory, error) {
	var categories []entities.FundraisingCategory
	if err := r.db.Find(&categories).Error; err != nil {
		return []entities.FundraisingCategory{}, err
	}
	return categories, nil
}

func (r *fundraisingRepository) FindByCategoryID(id int) ([]entities.Fundraising, error) {
	var fundraisings []entities.Fundraising
	if err := r.db.Preload("FundraisingCategory").Where("fundraising_category_id = ?", id).Find(&fundraisings).Error; err != nil {
		return []entities.Fundraising{}, err
	}
	return fundraisings, nil
}
