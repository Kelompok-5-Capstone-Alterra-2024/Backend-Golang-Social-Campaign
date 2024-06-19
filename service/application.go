package service

import (
	"capstone/entities"
	"capstone/repositories"
	"errors"
)

type ApplicationService interface {
	RegisterApplication(volunteerID uint, application entities.Application) (entities.Application, error)
	GetAllApplications(page, limit int) ([]entities.Application, int64, error)
	GetApplicationByID(id uint) (entities.Application, error)
	DeleteApplicationByID(id uint) error
	GetApplicationByVacancyID(vacancyID uint, page, limit int) ([]entities.Application, int, error)
}

type applicationService struct {
	applicationRepository repositories.ApplicationRepository
	volunteerRepo         repositories.VolunteerRepository
}

func NewApplicationService(applicationRepository repositories.ApplicationRepository, volunteerRepo repositories.VolunteerRepository) *applicationService {
	return &applicationService{applicationRepository, volunteerRepo}
}

func (s *applicationService) RegisterApplication(volunteerID uint, application entities.Application) (entities.Application, error) {
	volunteer, err := s.volunteerRepo.FindByID(volunteerID)
	if err != nil {
		return entities.Application{}, err
	}

	if application.VacancyID != volunteer.ID {
		return entities.Application{}, errors.New("volunteer opportunity doesn't match with the application")
	}

	// Check if application already exists for the given customer and vacancy
	existingApplication, err := s.applicationRepository.FindByCustomerIDAndVacancyID(application.UserID, application.VacancyID)
	if err == nil && existingApplication.ID != 0 {
		return entities.Application{}, errors.New("customer has already applied for this vacancy")
	}

	if volunteer.RegisteredVolunteer >= volunteer.TargetVolunteer {
		return entities.Application{}, errors.New("volunteer opportunity has reached its target")
	}

	volunteer.RegisteredVolunteer++
	if _, err := s.volunteerRepo.Update(volunteer); err != nil {
		return entities.Application{}, err
	}

	return s.applicationRepository.Save(application)
}

func (s *applicationService) GetAllApplications(page, limit int) ([]entities.Application, int64, error) {
	offset := (page - 1) * limit
	return s.applicationRepository.FindAll(offset, limit)
}

func (s *applicationService) GetApplicationByID(id uint) (entities.Application, error) {
	return s.applicationRepository.FindByID(id)
}

func (s *applicationService) DeleteApplicationByID(id uint) error {
	return s.applicationRepository.DeleteByID(id)
}

func (s *applicationService) GetApplicationByVacancyID(vacancyID uint, page, limit int) ([]entities.Application, int, error) {
	return s.applicationRepository.FindByVacancyID(vacancyID, page, limit)
}
