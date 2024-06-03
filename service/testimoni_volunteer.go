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
