package service_test

import (
	"capstone/entities"
	"capstone/mocks/repositories"
	"capstone/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestApplicationByID(t *testing.T) {
	testCases := []struct {
		title                   string
		mockApplicationbyID     entities.Application
		mockErrorbyID           error
		id                      uint
		expectedApplicationbyID entities.Application
		expectedErrorbyID       error
	}{
		{
			title: "Success Get Application by ID",
			mockApplicationbyID: entities.Application{
				Model: gorm.Model{
					ID: 1,
				},
				VacancyID:  1,
				UserID:     1,
				IgImageURL: "image.jpg",
				YtImageURL: "image.jpg",
				Job:        "Software Engineer",
				Reason:     "I want to helping my community",
			},
			mockErrorbyID: nil,
			expectedApplicationbyID: entities.Application{
				Model: gorm.Model{
					ID: 1,
				},
				VacancyID:  1,
				UserID:     1,
				IgImageURL: "image.jpg",
				YtImageURL: "image.jpg",
				Job:        "Software Engineer",
				Reason:     "I want to helping my community",
			},
			expectedErrorbyID: nil,
		},
		{
			title:                   "Failed Get Application by ID",
			mockApplicationbyID:     entities.Application{},
			mockErrorbyID:           errors.New("Failed to get application by ID"),
			expectedApplicationbyID: entities.Application{},
			expectedErrorbyID:       errors.New("Failed to get application by ID"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewApplicationRepository(t)

		repository.On("FindByID", testCase.id).Return(testCase.mockApplicationbyID, testCase.mockErrorbyID)

		service := service.NewApplicationService(repository, nil)
		application, err := service.GetApplicationByID(testCase.id)

		assert.Equal(t, testCase.expectedErrorbyID, err)
		assert.Equal(t, testCase.expectedApplicationbyID, application)
	}
}

func TestGetAllApplications(t *testing.T) {
	testCases := []struct {
		title                string
		mockApplications     []entities.Application
		mockError            error
		page                 int
		limit                int
		expectedApplications []entities.Application
		expectedError        error
	}{
		{
			title: "Success Get All Applications",
			mockApplications: []entities.Application{
				{
					Model: gorm.Model{
						ID: 1,
					},
					VacancyID:  1,
					UserID:     1,
					IgImageURL: "image.jpg",
					YtImageURL: "image.jpg",
					Job:        "Software Engineer",
					Reason:     "I want to helping my community",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					VacancyID:  2,
					UserID:     2,
					IgImageURL: "image.jpg",
					YtImageURL: "image.jpg",
					Job:        "Software Engineer",
					Reason:     "I want to helping my community",
				},
			},
			mockError: nil,
			page:      1,
			limit:     10,
			expectedApplications: []entities.Application{
				{
					Model: gorm.Model{
						ID: 1,
					},
					VacancyID:  1,
					UserID:     1,
					IgImageURL: "image.jpg",
					YtImageURL: "image.jpg",
					Job:        "Software Engineer",
					Reason:     "I want to helping my community",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					VacancyID:  2,
					UserID:     2,
					IgImageURL: "image.jpg",
					YtImageURL: "image.jpg",
					Job:        "Software Engineer",
					Reason:     "I want to helping my community",
				},
			},
			expectedError: nil,
		},
		{
			title:                "Failed Get All Applications",
			mockApplications:     []entities.Application{},
			mockError:            errors.New("Failed to get all applications"),
			page:                 1,
			limit:                10,
			expectedApplications: []entities.Application{},
			expectedError:        errors.New("Failed to get all applications"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewApplicationRepository(t)

		repository.On("FindAll", (testCase.page-1)*testCase.limit, testCase.limit).Return(testCase.mockApplications, int64(len(testCase.mockApplications)), testCase.mockError)

		service := service.NewApplicationService(repository, nil)
		applications, _, err := service.GetAllApplications(testCase.page, testCase.limit)

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedApplications, applications)
	}
}

func TestDeleteApplicationByID(t *testing.T) {
	testCases := []struct {
		title             string
		mockError         error
		id                uint
		expectedError     error
		expectedErrorCode int
	}{
		{
			title:             "Success Delete Application by ID",
			mockError:         nil,
			id:                1,
			expectedError:     nil,
			expectedErrorCode: 0,
		},
		{
			title:             "Failed Delete Application by ID",
			mockError:         errors.New("Failed to delete application by ID"),
			id:                1,
			expectedError:     errors.New("Failed to delete application by ID"),
			expectedErrorCode: 1,
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewApplicationRepository(t)

		repository.On("DeleteByID", testCase.id).Return(testCase.mockError)

		service := service.NewApplicationService(repository, nil)
		err := service.DeleteApplicationByID(testCase.id)

		assert.Equal(t, testCase.expectedError, err)
	}
}

// test RegisterApplication

// test GetApplicationByVacancyID
