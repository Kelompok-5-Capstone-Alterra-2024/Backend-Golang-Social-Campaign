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

func TestCreateTestimoniVolunteer(t *testing.T) {

	testCases := []struct {
		title                           string
		mockTestimoniVolunteer          entities.TestimoniVolunteer
		mockErrorTestimoniVolunteer     error
		expectedTestimoniVolunteer      entities.TestimoniVolunteer
		expectedErrorTestimoniVolunteer error
	}{
		{
			title: "Success Create Testimoni Volunteer",
			mockTestimoniVolunteer: entities.TestimoniVolunteer{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:    1,
				VacancyID: 1,
				Testimoni: "Testimoni",
				Rating:    "5",
			},
			mockErrorTestimoniVolunteer: nil,
			expectedTestimoniVolunteer: entities.TestimoniVolunteer{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:    1,
				VacancyID: 1,
				Testimoni: "Testimoni",
				Rating:    "5",
			},
			expectedErrorTestimoniVolunteer: nil,
		},
		{
			title: "Failed Create Testimoni Volunteer",
			mockTestimoniVolunteer: entities.TestimoniVolunteer{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:    1,
				VacancyID: 1,
				Testimoni: "Testimoni",
				Rating:    "5",
			},
			mockErrorTestimoniVolunteer:     errors.New("Failed to create testimoni volunteer"),
			expectedTestimoniVolunteer:      entities.TestimoniVolunteer{},
			expectedErrorTestimoniVolunteer: errors.New("Failed to create testimoni volunteer"),
		},
	}

	for _, tt := range testCases {
		repository := repositories.NewTestimoniVolunteerRepository(t)
		repository.On("Create", tt.mockTestimoniVolunteer).Return(tt.expectedTestimoniVolunteer, tt.mockErrorTestimoniVolunteer)

		service := service.NewTestimoniVolunteerService(repository)
		testimoniVolunteer, err := service.CreateTestimoniVolunteer(tt.mockTestimoniVolunteer)

		assert.Equal(t, tt.expectedErrorTestimoniVolunteer, err)
		assert.Equal(t, tt.expectedTestimoniVolunteer, testimoniVolunteer)
	}
}

func TestFindTestimoniByID(t *testing.T) {

	testCases := []struct {
		title                      string
		mockID                     uint
		mockTestimoniVolunteer     entities.TestimoniVolunteer
		mockError                  error
		expectedTestimoniVolunteer entities.TestimoniVolunteer
		expectedError              error
	}{
		{
			title:  "Success Find Testimoni Volunteer By ID",
			mockID: 1,
			mockTestimoniVolunteer: entities.TestimoniVolunteer{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:    1,
				VacancyID: 1,
				Testimoni: "Testimoni",
				Rating:    "5",
			},
			mockError: nil,
			expectedTestimoniVolunteer: entities.TestimoniVolunteer{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:    1,
				VacancyID: 1,
				Testimoni: "Testimoni",
				Rating:    "5",
			},
			expectedError: nil,
		},
		{
			title:                      "Failed Find Testimoni Volunteer By ID",
			mockID:                     1,
			mockTestimoniVolunteer:     entities.TestimoniVolunteer{},
			mockError:                  errors.New("Testimoni Volunteer not found"),
			expectedTestimoniVolunteer: entities.TestimoniVolunteer{},
			expectedError:              errors.New("Testimoni Volunteer not found"),
		},
	}

	for _, tt := range testCases {
		repository := repositories.NewTestimoniVolunteerRepository(t)
		repository.On("FindByID", tt.mockID).Return(tt.mockTestimoniVolunteer, tt.mockError)

		service := service.NewTestimoniVolunteerService(repository)
		testimoniVolunteer, err := service.FindByID(tt.mockID)

		assert.Equal(t, tt.expectedError, err)
		assert.Equal(t, tt.expectedTestimoniVolunteer, testimoniVolunteer)
	}
}

func TestFindAllTestimoni(t *testing.T) {

	testCases := []struct {
		title                           string
		mockPage                        int
		mockLimit                       int
		mockTestimoniVolunteers         []entities.TestimoniVolunteer
		mockTotalTestimoniVolunteer     int
		mockError                       error
		expectedTestimoniVolunteers     []entities.TestimoniVolunteer
		expectedTotalTestimoniVolunteer int
		expectedError                   error
	}{
		{
			title:     "Success Find All Testimoni Volunteer",
			mockPage:  1,
			mockLimit: 10,
			mockTestimoniVolunteers: []entities.TestimoniVolunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:    1,
					VacancyID: 1,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					UserID:    2,
					VacancyID: 2,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
			},
			mockTotalTestimoniVolunteer: 2,
			mockError:                   nil,
			expectedTestimoniVolunteers: []entities.TestimoniVolunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:    1,
					VacancyID: 1,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					UserID:    2,
					VacancyID: 2,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
			},
			expectedTotalTestimoniVolunteer: 2,
			expectedError:                   nil,
		},
		{
			title:                           "Failed Find All Testimoni Volunteer",
			mockPage:                        1,
			mockLimit:                       10,
			mockTestimoniVolunteers:         []entities.TestimoniVolunteer{},
			mockTotalTestimoniVolunteer:     0,
			mockError:                       errors.New("Failed to find testimoni volunteer"),
			expectedTestimoniVolunteers:     []entities.TestimoniVolunteer{},
			expectedTotalTestimoniVolunteer: 0,
			expectedError:                   errors.New("Failed to find testimoni volunteer"),
		},
	}

	for _, tt := range testCases {
		repository := repositories.NewTestimoniVolunteerRepository(t)
		repository.On("FindAll", tt.mockPage, tt.mockLimit).Return(tt.mockTestimoniVolunteers, tt.mockTotalTestimoniVolunteer, tt.mockError)

		service := service.NewTestimoniVolunteerService(repository)
		testimoniVolunteers, total, err := service.FindAll(tt.mockPage, tt.mockLimit)

		assert.Equal(t, tt.expectedError, err)
		assert.Equal(t, tt.expectedTestimoniVolunteers, testimoniVolunteers)
		assert.Equal(t, tt.expectedTotalTestimoniVolunteer, total)
	}
}

func TestDeleteTestimoniVolunteer(t *testing.T) {

	testCases := []struct {
		title                      string
		mockID                     uint
		mockError                  error
		expectedError              error
		expectedTestimoniVolunteer error
	}{
		{
			title:                      "Success Delete Testimoni Volunteer",
			mockID:                     1,
			mockError:                  nil,
			expectedError:              nil,
			expectedTestimoniVolunteer: nil,
		},
		{
			title:                      "Failed Delete Testimoni Volunteer",
			mockID:                     1,
			mockError:                  errors.New("Failed to delete testimoni volunteer"),
			expectedError:              errors.New("Failed to delete testimoni volunteer"),
			expectedTestimoniVolunteer: errors.New("Failed to delete testimoni volunteer"),
		},
	}

	for _, tt := range testCases {
		repository := repositories.NewTestimoniVolunteerRepository(t)
		repository.On("Delete", tt.mockID).Return(tt.mockError)

		service := service.NewTestimoniVolunteerService(repository)
		err := service.DeleteTestimoniVolunteer(tt.mockID)

		assert.Equal(t, tt.expectedError, err)
	}
}

func TestFindAllByVacancyID(t *testing.T) {

	testCases := []struct {
		title                       string
		mockVacancyID               uint
		mockTestimoniVolunteers     []entities.TestimoniVolunteer
		mockError                   error
		expectedTestimoniVolunteers []entities.TestimoniVolunteer
		expectedError               error
	}{
		{
			title:         "Success Find All By Vacancy ID",
			mockVacancyID: 1,
			mockTestimoniVolunteers: []entities.TestimoniVolunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:    1,
					VacancyID: 1,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					UserID:    2,
					VacancyID: 1,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
			},
			mockError: nil,
			expectedTestimoniVolunteers: []entities.TestimoniVolunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:    1,
					VacancyID: 1,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					UserID:    2,
					VacancyID: 1,
					Testimoni: "Testimoni",
					Rating:    "5",
				},
			},
			expectedError: nil,
		},
		{
			title:                       "Failed Find All By Vacancy ID",
			mockVacancyID:               1,
			mockTestimoniVolunteers:     []entities.TestimoniVolunteer{},
			mockError:                   errors.New("Failed to find testimoni volunteer"),
			expectedTestimoniVolunteers: []entities.TestimoniVolunteer{},
			expectedError:               errors.New("Failed to find testimoni volunteer"),
		},
	}

	for _, tt := range testCases {
		repository := repositories.NewTestimoniVolunteerRepository(t)
		repository.On("FindAllByVacancyID", tt.mockVacancyID).Return(tt.mockTestimoniVolunteers, tt.mockError)

		service := service.NewTestimoniVolunteerService(repository)
		testimoniVolunteers, err := service.FindAllByVacancyID(tt.mockVacancyID)

		assert.Equal(t, tt.expectedError, err)
		assert.Equal(t, tt.expectedTestimoniVolunteers, testimoniVolunteers)
	}
}

// func TestCustomerJoinedVolunteer(t *testing.T) {

// 	testCases := []struct {
// 		title                      string
// 		mockCustomerID             uint
// 		mockVolunteerID            uint
// 		mockJoined                 bool
// 		mockError                  error
// 		expectedJoined             bool
// 		expectedError              error
// 		expectedTestimoniVolunteer error
// 	}{
// 		{
// 			title:           "Success Customer Joined Volunteer",
// 			mockCustomerID:  1,
// 			mockVolunteerID: 1,
// 			mockJoined:      true,
// 			mockError:       nil,
// 			expectedJoined:  true,
// 			expectedError:   nil,
// 		},
// 		{
// 			title:           "Failed Customer Joined Volunteer",
// 			mockCustomerID:  1,
// 			mockVolunteerID: 1,
// 			mockJoined:      false,
// 			mockError:       errors.New("Failed to check customer joined volunteer"),
// 			expectedJoined:  false,
// 			expectedError:   errors.New("Failed to check customer joined volunteer"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewTestimoniVolunteerRepository(t)
// 		repository.On("CustomerJoinedVolunteer", tt.mockCustomerID, tt.mockVolunteerID).Return(tt.mockJoined, tt.mockError)

// 		service := service.NewTestimoniVolunteerService(repository)
// 		joined, err := service.CustomerJoinedVolunteer(tt.mockCustomerID, tt.mockVolunteerID)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedJoined, joined)
// 	}
// }

// func TestHasCustomerGivenTestimony(t *testing.T) {

// 	testCases := []struct {
// 		title                      string
// 		mockCustomerID             uint
// 		mockVolunteerID            uint
// 		mockHasGiven               bool
// 		mockError                  error
// 		expectedHasGiven           bool
// 		expectedError              error
// 		expectedTestimoniVolunteer error
// 	}{
// 		{
// 			title:            "Success Has Customer Given Testimony",
// 			mockCustomerID:   1,
// 			mockVolunteerID:  1,
// 			mockHasGiven:     true,
// 			mockError:        nil,
// 			expectedHasGiven: true,
// 			expectedError:    nil,
// 		},
// 		{
// 			title:            "Failed Has Customer Given Testimony",
// 			mockCustomerID:   1,
// 			mockVolunteerID:  1,
// 			mockHasGiven:     false,
// 			mockError:        errors.New("Failed to check customer given testimony"),
// 			expectedHasGiven: false,
// 			expectedError:    errors.New("Failed to check customer given testimony"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewTestimoniVolunteerRepository(t)
// 		repository.On("HasCustomerGivenTestimony", tt.mockCustomerID, tt.mockVolunteerID).Return(tt.mockHasGiven, tt.mockError)

// 		service := service.NewTestimoniVolunteerService(repository)
// 		hasGiven, err := service.HasCustomerGivenTestimony(tt.mockCustomerID, tt.mockVolunteerID)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedHasGiven, hasGiven)
// 	}
// }
