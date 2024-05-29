package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type VolunteerService interface {
	FindAll() ([]entities.Volunteer, error)
	FindByID(id uint) (entities.Volunteer, error)
	CreateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error)
}

type volunteerService struct {
	volunteerRepository repositories.VolunteerRepository
}

func NewVolunteerService(volunteerRepository repositories.VolunteerRepository) *volunteerService {
	return &volunteerService{volunteerRepository}
}

func (s *volunteerService) FindAll() ([]entities.Volunteer, error) {
	return s.volunteerRepository.FindAll()
}

func (s *volunteerService) FindByID(id uint) (entities.Volunteer, error) {
	return s.volunteerRepository.FindByID(id)
}

func (s *volunteerService) CreateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error) {
	return s.volunteerRepository.Save(volunteer)
}
