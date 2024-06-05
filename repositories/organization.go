package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Save(organization entities.Organization) (entities.Organization, error)
	FindAll() ([]entities.Organization, error)
	FindByID(id int) (entities.Organization, error)
	FindFundraisingByOrganizationID(id int) ([]entities.Fundraising, error)
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *organizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) Save(organization entities.Organization) (entities.Organization, error) {
	if err := r.db.Create(&organization).Error; err != nil {
		return entities.Organization{}, err
	}
	return organization, nil
}

func (r *organizationRepository) FindAll() ([]entities.Organization, error) {
	var organizations []entities.Organization
	if err := r.db.Find(&organizations).Error; err != nil {
		return []entities.Organization{}, err
	}
	return organizations, nil
}

func (r *organizationRepository) FindByID(id int) (entities.Organization, error) {
	var organization entities.Organization
	if err := r.db.Where("id = ?", id).First(&organization).Error; err != nil {
		return entities.Organization{}, err
	}
	return organization, nil
}

func (r *organizationRepository) FindFundraisingByOrganizationID(id int) ([]entities.Fundraising, error) {
	var fundraisings []entities.Fundraising
	if err := r.db.Preload("Organization").Preload("FundraisingCategory").Where("organization_id = ?", id).Find(&fundraisings).Error; err != nil {
		return []entities.Fundraising{}, err
	}
	return fundraisings, nil
}
