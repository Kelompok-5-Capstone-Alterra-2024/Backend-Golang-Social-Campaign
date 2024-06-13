package service

import (
	"capstone/dto"
	"capstone/entities"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"context"
	"fmt"
)

type AdminService interface {
	// Login(request dto.LoginRequest) (entities.Admin, error)
	Login(request dto.LoginRequest) (entities.Admin, string, string, error)
	GetFundraisings(limit int, offset int) ([]entities.Fundraising, error)
	CreateFudraising(ctx context.Context, fundraising entities.Fundraising) (entities.Fundraising, error)
	SaveImageFundraising(id uint, image string) (entities.Fundraising, error)
	DeleteFundraising(id uint) error
	UpdateFundraising(id uint, fundraising entities.Fundraising) (entities.Fundraising, error)
	GetFundraisingByID(id int) (entities.Fundraising, error)
	GetDonationByFundraisingID(id int, limit int, offset int) ([]entities.Donation, error)

	GetOrganizations(limit int, offset int) ([]entities.Organization, error)
	GetOrganizationByID(id int) (entities.Organization, error)
	UpdateOrganization(id uint, organization entities.Organization) (entities.Organization, error)
	DeleteOrganization(id uint) error
	GetVolunteerByOrganizationID(id int, page, limit int) ([]entities.Volunteer, int, error)
	GetFundraisingByOrganizationID(id int, page, limit int) ([]entities.Fundraising, int, error)
	SaveImageOraganization(id uint, image string) (entities.Organization, error)

	GetUsers(limit int, offset int) ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	UpdateUserByID(id uint, user entities.User) (entities.User, error)
	GetDonationsByUserID(id int, page int, limit int) ([]entities.DonationManual, int, error)
	GetVolunteersByUserID(id int, page int, limit int) ([]entities.Application, int, error)
	DeleteUserWithDonations(id uint) error

	GetAllDonations(page, limit int) ([]entities.DonationManual, int, error)
	AddAmountToUserDonation(id uint, amount int) (entities.DonationManual, error)
}

type adminService struct {
	adminRepository repositories.AdminRepository
	userRepository  repositories.UserRepository
}

func NewAdminService(adminRepository repositories.AdminRepository, userRepository repositories.UserRepository) *adminService {
	return &adminService{adminRepository, userRepository}
}

// func (s *adminService) Login(request dto.LoginRequest) (entities.Admin, error) {
// 	username := request.Username
// 	password := request.Password

// 	admin, err := s.adminRepository.FindByUsername(username)
// 	if err != nil {
// 		return admin, err
// 	}

// 	if admin.Password != password {
// 		return admin, errors.New("wrong password")
// 	}

// 	admin.Token = middleware.GenerateToken(admin.ID, admin.Username, "admin")

// 	return admin, nil
// }

func (s *adminService) Login(request dto.LoginRequest) (entities.Admin, string, string, error) {
	username := request.Username
	password := request.Password

	admin, err := s.adminRepository.FindByUsername(username)
	if err != nil {
		return admin, "", "", err
	}

	if admin.Password != password {
		return admin, "", "", fmt.Errorf("wrong password")
	}

	accessToken, refreshToken := middleware.GenerateToken(admin.ID, admin.Username, "admin")
	admin.Token = accessToken

	return admin, accessToken, refreshToken, nil
}

func (s *adminService) GetFundraisings(limit int, offset int) ([]entities.Fundraising, error) {
	return s.adminRepository.FindAllFundraising(limit, offset)
}

func (s *adminService) CreateFudraising(ctx context.Context, fundraising entities.Fundraising) (entities.Fundraising, error) {
	return s.adminRepository.Create(fundraising)
}

func (s *adminService) SaveImageFundraising(id uint, image string) (entities.Fundraising, error) {
	fund, err := s.adminRepository.FindFundraisingByID(int(id))

	if err != nil {
		return fund, err
	}

	fund.ImageUrl = image
	updatedFund, err := s.adminRepository.UpdateFundraisingByID(id, fund)

	if err != nil {
		return updatedFund, err
	}
	return updatedFund, nil
}

func (s *adminService) DeleteFundraising(id uint) error {
	return s.adminRepository.DeleteFundraising(id)
}

func (s *adminService) UpdateFundraising(id uint, fundraising entities.Fundraising) (entities.Fundraising, error) {
	return s.adminRepository.UpdateFundraisingByID(id, fundraising)
}

func (s *adminService) GetFundraisingByID(id int) (entities.Fundraising, error) {
	return s.adminRepository.FindFundraisingByID(id)
}

func (s *adminService) GetDonationByFundraisingID(id int, limit int, offset int) ([]entities.Donation, error) {
	return s.adminRepository.FindDonationsByFundraisingID(id, limit, offset)
}

func (s *adminService) GetOrganizations(limit int, offset int) ([]entities.Organization, error) {
	return s.adminRepository.FindOrganizations(limit, offset)
}

func (s *adminService) GetOrganizationByID(id int) (entities.Organization, error) {
	return s.adminRepository.FindOrganizationByID(id)
}

func (s *adminService) UpdateOrganization(id uint, organization entities.Organization) (entities.Organization, error) {
	return s.adminRepository.UpdateOrganizationByID(id, organization)
}

func (s *adminService) DeleteOrganization(id uint) error {
	return s.adminRepository.DeleteOrganizationByID(id)
}

func (s *adminService) GetVolunteerByOrganizationID(id int, page, limit int) ([]entities.Volunteer, int, error) {
	return s.adminRepository.GetVolunteerByOrganizationID(id, page, limit)
}

func (s *adminService) GetFundraisingByOrganizationID(id int, page, limit int) ([]entities.Fundraising, int, error) {
	return s.adminRepository.GetFundraisingByOrganizationID(id, page, limit)
}

func (s *adminService) SaveImageOraganization(id uint, image string) (entities.Organization, error) {
	org, err := s.adminRepository.FindOrganizationByID(int(id))

	if err != nil {
		return org, err
	}

	org.Avatar = image
	updatedOrg, err := s.adminRepository.UpdateOrganizationByID(id, org)

	if err != nil {
		return updatedOrg, err
	}

	return updatedOrg, nil
}

func (s *adminService) GetUsers(limit int, offset int) ([]entities.User, error) {
	return s.adminRepository.FindUsers(limit, offset)
}

func (s *adminService) GetUserByID(id int) (entities.User, error) {
	return s.userRepository.FindByID(uint(id))
}

func (s *adminService) UpdateUserByID(id uint, user entities.User) (entities.User, error) {
	return s.adminRepository.UpdateUserByID(id, user)
}

func (s *adminService) GetDonationsByUserID(id int, page int, limit int) ([]entities.DonationManual, int, error) {

	return s.adminRepository.FindDonationsByUserID(id, page, limit)

}

func (s *adminService) GetVolunteersByUserID(id int, page int, limit int) ([]entities.Application, int, error) {

	return s.adminRepository.FindVolunteersByUserID(id, page, limit)
}

func (s *adminService) DeleteUserWithDonations(id uint) error {
	return s.adminRepository.DeleteUserWithDonations(id)
}

func (s *adminService) GetAllDonations(page, limt int) ([]entities.DonationManual, int, error) {
	return s.adminRepository.FindAllDonations(page, limt)
}

func (s *adminService) AddAmountToUserDonation(id uint, amount int) (entities.DonationManual, error) {
	fundraising, err := s.adminRepository.FindFundraisingByDonationID(int(id))
	if err != nil {
		return entities.DonationManual{}, err
	}
	donation, err := s.adminRepository.AddAmountToUserDonation(id, amount)

	if err != nil {
		return entities.DonationManual{}, err
	}

	fundraising.CurrentProgress += donation.Amount
	if fundraising.CurrentProgress == fundraising.GoalAmount {
		fundraising.Status = "Achived"
	}
	_, err = s.adminRepository.UpdateFundraisingByID(fundraising.ID, fundraising)
	if err != nil {
		return entities.DonationManual{}, err
	}

	return entities.DonationManual{}, nil

}
