package service

import (
	"capstone/entities"
	"capstone/repositories"
	"context"
	"errors"
)

type FundraisingService interface {
	FindFundraisings(ctx context.Context, limit int, offset int) ([]entities.Fundraising, error)
	FindTopFundraisings() ([]entities.Fundraising, error)
	FindFundraisingByID(id int) (entities.Fundraising, error)
	FindAllFundraisingCategories() ([]entities.FundraisingCategory, error)
	FindFundraisingByCategoryID(id int, limit int, offset int) ([]entities.Fundraising, error)
}

type fundraisingService struct {
	fundraisingRepository repositories.FundraisingRepository
}

func NewFundraisingService(fundraisingRepository repositories.FundraisingRepository) *fundraisingService {
	return &fundraisingService{fundraisingRepository}
}

func (s *fundraisingService) FindFundraisings(ctx context.Context, limit int, offset int) ([]entities.Fundraising, error) {
	return s.fundraisingRepository.FindAll(limit, offset)
}

func (s *fundraisingService) FindTopFundraisings() ([]entities.Fundraising, error) {
	return s.fundraisingRepository.FindTopFundraisings()
}

func (s *fundraisingService) FindFundraisingByID(id int) (entities.Fundraising, error) {
	fundraising, err := s.fundraisingRepository.FindByID(id)

	if err != nil {
		return entities.Fundraising{}, errors.New("fundraising not found")
	}

	return fundraising, nil
}

func (s *fundraisingService) FindAllFundraisingCategories() ([]entities.FundraisingCategory, error) {
	return s.fundraisingRepository.FindAllCategories()
}

func (s *fundraisingService) FindFundraisingByCategoryID(id int, limit int, offset int) ([]entities.Fundraising, error) {
	return s.fundraisingRepository.FindByCategoryID(id, limit, offset)
}
