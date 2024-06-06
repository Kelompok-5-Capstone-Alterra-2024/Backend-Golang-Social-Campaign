package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (entities.Admin, error)
	FindAllFundraising(limit int, offset int) ([]entities.Fundraising, error)
	Create(fundraising entities.Fundraising) (entities.Fundraising, error)
	DeleteFundraising(id uint) error
	UpdateFundraisingByID(id uint, fundraising entities.Fundraising) (entities.Fundraising, error)
	FindFundraisingByID(id int) (entities.Fundraising, error)
	FindDonationsByFundraisingID(id int, limit int, offset int) ([]entities.Donation, error)

	FindOrganizations(limit int, offset int) ([]entities.Organization, error)
	FindOrganizationByID(id int) (entities.Organization, error)
	UpdateOrganizationByID(id uint, organization entities.Organization) (entities.Organization, error)
	DeleteOrganizationByID(id uint) error

	FindUsers(limit int, offset int) ([]entities.User, error)
	FindDonationsByUserID(id int, limit int, offset int) ([]entities.Donation, error)
	DeleteUserWithDonations(id uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByUsername(username string) (entities.Admin, error) {
	var admin entities.Admin
	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *adminRepository) FindAllFundraising(limit int, offset int) ([]entities.Fundraising, error) {
	var fundraisings []entities.Fundraising
	if err := r.db.Preload("Organization").Limit(limit).Offset(offset).Find(&fundraisings).Error; err != nil {
		return []entities.Fundraising{}, err
	}
	return fundraisings, nil
}

func (r *adminRepository) Create(fundraising entities.Fundraising) (entities.Fundraising, error) {
	if err := r.db.Create(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) UpdateFundraisingByID(id uint, fundraising entities.Fundraising) (entities.Fundraising, error) {
	if err := r.db.Model(&fundraising).Where("id = ?", id).Updates(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) DeleteFundraising(id uint) error {
	if err := r.db.Delete(&entities.Fundraising{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) FindFundraisingByID(id int) (entities.Fundraising, error) {
	var fundraising entities.Fundraising
	if err := r.db.Where("id = ?", id).First(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) FindDonationsByFundraisingID(id int, limit int, offset int) ([]entities.Donation, error) {
	var donations []entities.Donation
	if err := r.db.Preload("User").Preload("Fundraising").Where("fundraising_id = ?", id).Limit(limit).Offset(offset).Find(&donations).Error; err != nil {
		return []entities.Donation{}, err
	}
	return donations, nil
}

func (r *adminRepository) FindOrganizations(limit int, offset int) ([]entities.Organization, error) {
	var organizations []entities.Organization
	if err := r.db.Limit(limit).Offset(offset).Find(&organizations).Error; err != nil {
		return []entities.Organization{}, err
	}
	return organizations, nil
}

func (r *adminRepository) FindOrganizationByID(id int) (entities.Organization, error) {
	var organization entities.Organization
	if err := r.db.Where("id = ?", id).First(&organization).Error; err != nil {
		return entities.Organization{}, err
	}
	return organization, nil
}

func (r *adminRepository) UpdateOrganizationByID(id uint, organization entities.Organization) (entities.Organization, error) {
	if err := r.db.Model(&organization).Where("id = ?", id).Updates(&organization).Error; err != nil {
		return entities.Organization{}, err
	}
	return organization, nil
}

func (r *adminRepository) DeleteOrganizationByID(id uint) error {
	if err := r.db.Delete(&entities.Organization{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) FindUsers(limit int, offset int) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return []entities.User{}, err
	}
	return users, nil
}

func (r *adminRepository) FindDonationsByUserID(id int, limit int, offset int) ([]entities.Donation, error) {
	var donations []entities.Donation
	if err := r.db.Preload("Fundraising.Organization").Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&donations).Error; err != nil {
		return []entities.Donation{}, err
	}
	return donations, nil
}

func (r *adminRepository) DeleteUserWithDonations(id uint) error {
	if err := r.db.Where("user_id = ?", id).Delete(&entities.Donation{}).Error; err != nil {
		return err
	}

	// Hapus user
	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
