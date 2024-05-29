package service

import (
	"capstone/entities"
	"capstone/repositories"
)

type ApplicationService interface {
	RegisterApplication(application entities.Application) (entities.Application, error)
}

type applicationService struct {
	applicationRepository repositories.ApplicationRepository
}

func NewApplicationService(applicationRepository repositories.ApplicationRepository) *applicationService {
	return &applicationService{applicationRepository}
}

func (s *applicationService) RegisterApplication(application entities.Application) (entities.Application, error) {
	return s.applicationRepository.Save(application)
}
