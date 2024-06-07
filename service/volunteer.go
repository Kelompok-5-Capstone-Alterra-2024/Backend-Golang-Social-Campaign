package service

import (
	"capstone/entities"
	"capstone/repositories"
	"errors"
	"time"
)

type VolunteerService interface {
	CreateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error)
	FindByID(id uint) (entities.Volunteer, error)
	FindAll(page int, limit int) ([]entities.Volunteer, int, error)
	ApplyForVolunteer(volunteerID uint, customerID uint) (entities.Volunteer, error)
	UpdateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error)
	UpdateVolunteerByID(id uint, volunteer entities.Volunteer) (entities.Volunteer, error)
	DeleteVolunteer(id uint) error
}

type volunteerService struct {
	volunteerRepository repositories.VolunteerRepository
}

func NewVolunteerService(volunteerRepo repositories.VolunteerRepository) VolunteerService {
	return &volunteerService{volunteerRepository: volunteerRepo}
}

func (s *volunteerService) CreateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error) {
	return s.volunteerRepository.Create(volunteer)
}

func (s *volunteerService) FindByID(id uint) (entities.Volunteer, error) {
	return s.volunteerRepository.FindByID(id)
}

func (s *volunteerService) FindAll(page int, limit int) ([]entities.Volunteer, int, error) {
	return s.volunteerRepository.FindAll(page, limit)
}

func (s *volunteerService) ApplyForVolunteer(volunteerID uint, customerID uint) (entities.Volunteer, error) {
	volunteer, err := s.volunteerRepository.FindByID(volunteerID)
	if err != nil {
		return volunteer, err
	}

	// Check if the customer has already applied for this volunteer
	application, err := s.volunteerRepository.FindApplicationByVolunteerAndCustomer(volunteerID, customerID)
	if err == nil && application.ID != 0 {
		return volunteer, errors.New("customer has already applied for this volunteer opportunity")
	}

	// Check if the volunteer opportunity has reached its target
	if volunteer.RegisteredVolunteer >= volunteer.TargetVolunteer {
		return volunteer, errors.New("volunteer opportunity has reached its target")
	}

	// Check if the current date is past the registration deadline
	loc, _ := time.LoadLocation("Asia/Jakarta")
	currentDate := time.Now().In(loc).Format("02/01/2006")
	registrationDeadline := volunteer.RegistrationDeadline.Format("02/01/2006")
	if currentDate == registrationDeadline {
		return volunteer, errors.New("registration deadline has passed")
	}

	// Register the customer
	volunteer.RegisteredVolunteer++
	return s.volunteerRepository.Update(volunteer)
}

func (s *volunteerService) UpdateVolunteer(volunteer entities.Volunteer) (entities.Volunteer, error) {
	return s.volunteerRepository.Update(volunteer)
}

func (s *volunteerService) UpdateVolunteerByID(id uint, volunteer entities.Volunteer) (entities.Volunteer, error) {
	return s.volunteerRepository.UpdateByID(id, volunteer)
}

func (s *volunteerService) DeleteVolunteer(id uint) error {
	return s.volunteerRepository.Delete(id)
}
