package service

import (
	"capstone/dto"
	"capstone/entities"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"errors"
)

type AdminService interface {
	Login(request dto.LoginRequest) (entities.Admin, error)
}

type adminService struct {
	adminRepository repositories.AdminRepository
}

func NewAdminService(adminRepository repositories.AdminRepository) *adminService {
	return &adminService{adminRepository}
}

func (s *adminService) Login(request dto.LoginRequest) (entities.Admin, error) {
	username := request.Username
	password := request.Password

	admin, err := s.adminRepository.FindByUsername(username)
	if err != nil {
		return admin, err
	}

	if admin.Password != password {
		return admin, errors.New("wrong password")
	}

	admin.Token = middleware.GenerateToken(admin.ID, admin.Username, "admin")

	return admin, nil
}
