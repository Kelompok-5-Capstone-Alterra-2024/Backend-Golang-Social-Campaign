package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type VolunteerService interface {
	FindAll(page, limit int) ([]entities.Volunteer, int64, error)
	FindByID(id uint) (entities.Volunteer, error)
	CreateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error)
}

type volunteerService struct {
	volunteerRepository repositories.VolunteerRepository
}

func NewVolunteerService(volunteerRepository repositories.VolunteerRepository) *volunteerService {
	return &volunteerService{volunteerRepository}
}

func (s *volunteerService) FindAll(page, limit int) ([]entities.Volunteer, int64, error) {
	return s.volunteerRepository.FindAll(page, limit)
}

func (s *volunteerService) FindByID(id uint) (entities.Volunteer, error) {
	return s.volunteerRepository.FindByID(id)
}

func (s *volunteerService) CreateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error) {
	return s.volunteerRepository.Save(volunteer)
}
