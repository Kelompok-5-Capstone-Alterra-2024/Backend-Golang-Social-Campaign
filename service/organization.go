package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type OrganizationService interface {
	CreateOrganization(organization entities.Organization) (entities.Organization, error)
	FindOrganizations() ([]entities.Organization, error)
	FindOrganizationByID(id int) (entities.Organization, error)
	FindFundraisingByOrganizationID(id int, limit int, offset int) ([]entities.Fundraising, error)
	FindVolunteersByOrganizationID(id int, limit int, offset int) ([]entities.Volunteer, error)
}

type organizationService struct {
	organizationRepository repositories.OrganizationRepository
}

func NewOrganizationService(organizationRepository repositories.OrganizationRepository) *organizationService {
	return &organizationService{organizationRepository}
}

func (s *organizationService) CreateOrganization(organization entities.Organization) (entities.Organization, error) {

	return s.organizationRepository.Save(organization)
}

func (s *organizationService) FindOrganizations() ([]entities.Organization, error) {
	return s.organizationRepository.FindAll()
}

func (s *organizationService) FindOrganizationByID(id int) (entities.Organization, error) {
	return s.organizationRepository.FindByID(id)
}

func (s *organizationService) FindFundraisingByOrganizationID(id int, limit int, offset int) ([]entities.Fundraising, error) {
	return s.organizationRepository.FindFundraisingByOrganizationID(id, limit, offset)
}

func (s *organizationService) FindVolunteersByOrganizationID(id int, limit int, offset int) ([]entities.Volunteer, error) {
	return s.organizationRepository.FindVolunteersByOrganizationID(id, limit, offset)
}
