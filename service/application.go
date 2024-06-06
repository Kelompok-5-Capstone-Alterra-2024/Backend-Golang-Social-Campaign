package service

import (
	"capstone/entities"
	"capstone/repositories"
	"errors"
)

type ApplicationService interface {
	RegisterApplication(application entities.Application) (entities.Application, error)
	GetAllApplications(page, limit int) ([]entities.Application, int64, error)
	GetApplicationByID(id uint) (entities.Application, error)
	DeleteApplicationByID(id uint) error
	GetApplicationByVacancyID(vacancyID uint) ([]entities.Application, error)
}

type applicationService struct {
	applicationRepository repositories.ApplicationRepository
}

func NewApplicationService(applicationRepository repositories.ApplicationRepository) *applicationService {
	return &applicationService{applicationRepository}
}

func (s *applicationService) RegisterApplication(application entities.Application) (entities.Application, error) {
	// Check if application already exists for the given customer and vacancy
	existingApplication, err := s.applicationRepository.FindByCustomerIDAndVacancyID(application.UserID, application.VacancyID)
	if err == nil && existingApplication.ID != 0 {
		return entities.Application{}, errors.New("customer has already applied for this vacancy")
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

func (s *applicationService) GetApplicationByVacancyID(vacancyID uint) ([]entities.Application, error) {
	return s.applicationRepository.FindByVacancyID(vacancyID)
}
