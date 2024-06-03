package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type OrganizationService interface {
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

func (s *organizationService) FindOrganizations() ([]entities.Organization, error) {
	return s.organizationRepository.FindAll()
}

func (s *organizationService) FindOrganizationByID(id int) (entities.Organization, error) {
	return s.organizationRepository.FindByID(id)
}

func (s *organizationService) FindFundraisingByOrganizationID(id int) ([]entities.Fundraising, error) {
	return s.organizationRepository.FindFundraisingByOrganizationID(id)
}
