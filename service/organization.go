package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/repositories"
)

type OrganizationService interface {
	CreateOrganization(input dto.OrganizationRequest) (entities.Organization, error)
	FindOrganizations() ([]entities.Organization, error)
	FindOrganizationByID(id int) (entities.Organization, error)
	FindFundraisingByOrganizationID(id int) ([]entities.Fundraising, error)
}

type organizationService struct {
	organizationRepository repositories.OrganizationRepository
}

func NewOrganizationService(organizationRepository repositories.OrganizationRepository) *organizationService {
	return &organizationService{organizationRepository}
}

func (s *organizationService) CreateOrganization(input dto.OrganizationRequest) (entities.Organization, error) {
	organization := entities.Organization{
		Name:        input.Name,
		Avatar:      input.Avatar,
		Description: input.Description,
		IsVerified:  input.IsVerified,
	}

	return s.organizationRepository.Save(organization)
}

func (s *organizationService) FindOrganizations() ([]entities.Organization, error) {
	return s.organizationRepository.FindAll()
}

func (s *organizationService) FindOrganizationByID(id int) (entities.Organization, error) {
	return s.organizationRepository.FindByID(id)
}

func (s *organizationService) FindFundraisingByOrganizationID(id int) ([]entities.Fundraising, error) {
	return s.organizationRepository.FindFundraisingByOrganizationID(id)
}
