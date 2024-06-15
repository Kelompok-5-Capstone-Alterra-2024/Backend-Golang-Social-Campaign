package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type TestimoniVolunteerService interface {
	CreateTestimoniVolunteer(testimoniVolunteer entities.TestimoniVolunteer) (entities.TestimoniVolunteer, error)
	FindByID(id uint) (entities.TestimoniVolunteer, error)
	FindAll(page, limit int) ([]entities.TestimoniVolunteer, int, error)
	DeleteTestimoniVolunteer(id uint) error
	FindAllByVacancyID(volunteerID uint, limit, offest int) ([]entities.TestimoniVolunteer, error)
	CustomerJoinedVolunteer(customerID, volunteerID uint) bool
	HasCustomerGivenTestimony(customerID, volunteerID uint) bool
}

type testimoniVolunteerService struct {
	repo repositories.TestimoniVolunteerRepository
}

func NewTestimoniVolunteerService(repo repositories.TestimoniVolunteerRepository) TestimoniVolunteerService {
	return &testimoniVolunteerService{repo: repo}
}

func (s *testimoniVolunteerService) CreateTestimoniVolunteer(testimoniVolunteer entities.TestimoniVolunteer) (entities.TestimoniVolunteer, error) {
	return s.repo.Create(testimoniVolunteer)
}

func (s *testimoniVolunteerService) FindByID(id uint) (entities.TestimoniVolunteer, error) {
	return s.repo.FindByID(id)
}

func (s *testimoniVolunteerService) FindAll(page, limit int) ([]entities.TestimoniVolunteer, int, error) {
	return s.repo.FindAll(page, limit)
}

func (s *testimoniVolunteerService) DeleteTestimoniVolunteer(id uint) error {
	return s.repo.Delete(id)
}

func (s *testimoniVolunteerService) FindAllByVacancyID(volunteerID uint, limit, offest int) ([]entities.TestimoniVolunteer, error) {
	return s.repo.FindAllByVacancyID(volunteerID, limit, offest)
}

func (s *testimoniVolunteerService) CustomerJoinedVolunteer(customerID, volunteerID uint) bool {
	joined, err := s.repo.CustomerJoinedVolunteer(customerID, volunteerID)
	if err != nil {
		return false
	}
	return joined
}

func (s *testimoniVolunteerService) HasCustomerGivenTestimony(customerID, volunteerID uint) bool {
	hasGiven, err := s.repo.HasCustomerGivenTestimony(customerID, volunteerID)
	if err != nil {
		return false
	}
	return hasGiven
}
