package service_test

import (
	"capstone/entities"
	"capstone/mocks/repositories"
	"capstone/service"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateVolunteer(t *testing.T) {
	testCases := []struct {
		title             string
		mockVolunteer     entities.Volunteer
		mockError         error
		expectedVolunteer entities.Volunteer
		expectedError     error
	}{
		{
			title: "Success Create Volunteer",
			mockVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			mockError: nil,
			expectedVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			expectedError: nil,
		},
		{
			title:             "Failed Create Volunteer",
			mockVolunteer:     entities.Volunteer{},
			mockError:         errors.New("Unexpected Error"),
			expectedVolunteer: entities.Volunteer{},
			expectedError:     errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewVolunteerRepository(t)
		repository.On("Create", tc.mockVolunteer).Return(tc.expectedVolunteer, tc.mockError)

		s := service.NewVolunteerService(repository)

		volunteer, err := s.CreateVolunteer(tc.mockVolunteer)

		assert.Equal(t, tc.expectedVolunteer, volunteer)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestUpdateVolunteer(t *testing.T) {
	testCases := []struct {
		title             string
		mockVolunteer     entities.Volunteer
		mockError         error
		expectedVolunteer entities.Volunteer
		expectedError     error
	}{
		{
			title: "Success Update Volunteer",
			mockVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			mockError: nil,
			expectedVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			expectedError: nil,
		},
		{
			title:             "Failed Update Volunteer",
			mockVolunteer:     entities.Volunteer{},
			mockError:         errors.New("Unexpected Error"),
			expectedVolunteer: entities.Volunteer{},
			expectedError:     errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewVolunteerRepository(t)
		repository.On("Update", tc.mockVolunteer).Return(tc.expectedVolunteer, tc.mockError)

		s := service.NewVolunteerService(repository)

		volunteer, err := s.UpdateVolunteer(tc.mockVolunteer)

		assert.Equal(t, tc.expectedVolunteer, volunteer)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestFindVolunteerByID(t *testing.T) {
	testCases := []struct {
		title             string
		mockID            uint
		mockVolunteer     entities.Volunteer
		mockError         error
		expectedVolunteer entities.Volunteer
		expectedError     error
	}{
		{
			title:  "Success Find Volunteer By ID",
			mockID: 1,
			mockVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			mockError: nil,
			expectedVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			expectedError: nil,
		},
		{
			title:             "Failed Find Volunteer By ID",
			mockID:            1,
			mockVolunteer:     entities.Volunteer{},
			mockError:         errors.New("Unexpected Error"),
			expectedVolunteer: entities.Volunteer{},
			expectedError:     errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewVolunteerRepository(t)
		repository.On("FindByID", tc.mockID).Return(tc.expectedVolunteer, tc.mockError)

		s := service.NewVolunteerService(repository)

		volunteer, err := s.FindByID(tc.mockID)

		assert.Equal(t, tc.expectedVolunteer, volunteer)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestFindAllVolunteer(t *testing.T) {
	testCases := []struct {
		title              string
		mockPage           int
		mockLimit          int
		mockVolunteers     []entities.Volunteer
		mockTotal          int
		mockError          error
		expectedVolunteers []entities.Volunteer
		expectedTotal      int
		expectedError      error
	}{
		{
			title:     "Success Find All Volunteers",
			mockPage:  1,
			mockLimit: 10,
			mockVolunteers: []entities.Volunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:          "Helping Cleaning Bandung from Trash",
					OrganizationID: 1,
					Organization: entities.Organization{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Bandung Lover",
					},
					ContentActivity:      "We will clean the trash in Bandung",
					Location:             "Bandung",
					LinkWA:               "https://wa.me/628123456789",
					StartDate:            time.Now(),
					EndDate:              time.Now().AddDate(0, 0, 7),
					TargetVolunteer:      10,
					Status:               "active",
					RegisteredVolunteer:  0,
					RegistrationDeadline: time.Now().AddDate(0, 0, 5),
					ImageURL:             "https://imageurl.com",
				},
			},
			mockTotal: 1,
			mockError: nil,
			expectedVolunteers: []entities.Volunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:          "Helping Cleaning Bandung from Trash",
					OrganizationID: 1,
					Organization: entities.Organization{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Bandung Lover",
					},
					ContentActivity:      "We will clean the trash in Bandung",
					Location:             "Bandung",
					LinkWA:               "https://wa.me/628123456789",
					StartDate:            time.Now(),
					EndDate:              time.Now().AddDate(0, 0, 7),
					TargetVolunteer:      10,
					Status:               "active",
					RegisteredVolunteer:  0,
					RegistrationDeadline: time.Now().AddDate(0, 0, 5),
					ImageURL:             "https://imageurl.com",
				},
			},
			expectedTotal: 1,
			expectedError: nil,
		},
		{
			title:              "Failed Find All Volunteers",
			mockPage:           1,
			mockLimit:          10,
			mockVolunteers:     []entities.Volunteer{},
			mockTotal:          0,
			mockError:          errors.New("Unexpected Error"),
			expectedVolunteers: []entities.Volunteer{},
			expectedTotal:      0,
			expectedError:      errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewVolunteerRepository(t)
		repository.On("FindAll", tc.mockPage, tc.mockLimit).Return(tc.mockVolunteers, tc.mockTotal, tc.mockError)

		s := service.NewVolunteerService(repository)

		volunteers, total, err := s.FindAll(tc.mockPage, tc.mockLimit)

		assert.Equal(t, tc.expectedVolunteers, volunteers)
		assert.Equal(t, tc.expectedTotal, total)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestFindTopVolunteers(t *testing.T) {
	testCases := []struct {
		title              string
		mockVolunteers     []entities.Volunteer
		mockError          error
		expectedVolunteers []entities.Volunteer
		expectedError      error
	}{
		{
			title: "Success Find Top Volunteers",
			mockVolunteers: []entities.Volunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:          "Helping Cleaning Bandung from Trash",
					OrganizationID: 1,
					Organization: entities.Organization{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Bandung Lover",
					},
					ContentActivity:      "We will clean the trash in Bandung",
					Location:             "Bandung",
					LinkWA:               "https://wa.me/628123456789",
					StartDate:            time.Now(),
					EndDate:              time.Now().AddDate(0, 0, 7),
					TargetVolunteer:      10,
					Status:               "active",
					RegisteredVolunteer:  0,
					RegistrationDeadline: time.Now().AddDate(0, 0, 5),
					ImageURL:             "https://imageurl.com",
				},
			},
			mockError: nil,
			expectedVolunteers: []entities.Volunteer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:          "Helping Cleaning Bandung from Trash",
					OrganizationID: 1,
					Organization: entities.Organization{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Bandung Lover",
					},
					ContentActivity:      "We will clean the trash in Bandung",
					Location:             "Bandung",
					LinkWA:               "https://wa.me/628123456789",
					StartDate:            time.Now(),
					EndDate:              time.Now().AddDate(0, 0, 7),
					TargetVolunteer:      10,
					Status:               "active",
					RegisteredVolunteer:  0,
					RegistrationDeadline: time.Now().AddDate(0, 0, 5),
					ImageURL:             "https://imageurl.com",
				},
			},
			expectedError: nil,
		},
		{
			title:              "Failed Find Top Volunteers",
			mockVolunteers:     []entities.Volunteer{},
			mockError:          errors.New("Unexpected Error"),
			expectedVolunteers: []entities.Volunteer{},
			expectedError:      errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewVolunteerRepository(t)
		repository.On("FindTop").Return(tc.mockVolunteers, tc.mockError)

		s := service.NewVolunteerService(repository)

		volunteers, err := s.FindTopVolunteers()

		assert.Equal(t, tc.expectedVolunteers, volunteers)
		assert.Equal(t, tc.expectedError, err)
	}
}

// func TestApplyForVolunteer(t *testing.T) {
// 	testCases := []struct {
// 		title             string
// 		mockVolunteerID   uint
// 		mockCustomerID    uint
// 		mockVolunteer     entities.Volunteer
// 		mockApplication   entities.Application
// 		mockError         error
// 		expectedVolunteer entities.Volunteer
// 		expectedError     error
// 	}{
// 		{
// 			title:           "Success Apply For Volunteer",
// 			mockVolunteerID: 1,
// 			mockCustomerID:  1,
// 			mockVolunteer: entities.Volunteer{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Title:          "Helping Cleaning Bandung from Trash",
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Bandung Lover",
// 				},
// 				ContentActivity:      "We will clean the trash in Bandung",
// 				Location:             "Bandung",
// 				LinkWA:               "https://wa.me/628123456789",
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().AddDate(0, 0, 7),
// 				TargetVolunteer:      10,
// 				Status:               "active",
// 				RegisteredVolunteer:  0,
// 				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
// 				ImageURL:             "https://imageurl.com",
// 			},
// 			mockApplication: entities.Application{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				VacancyID: 1,
// 				UserID:    1,
// 			},
// 			mockError: nil,
// 			expectedVolunteer: entities.Volunteer{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Title:          "Helping Cleaning Bandung from Trash",
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Bandung Lover",
// 				},
// 				ContentActivity:      "We will clean the trash in Bandung",
// 				Location:             "Bandung",
// 				LinkWA:               "https://wa.me/628123456789",
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().AddDate(0, 0, 7),
// 				TargetVolunteer:      10,
// 				Status:               "active",
// 				RegisteredVolunteer:  1,
// 				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
// 				ImageURL:             "https://imageurl.com",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:           "Failed Apply For Volunteer",
// 			mockVolunteerID: 1,
// 			mockCustomerID:  1,
// 			mockVolunteer: entities.Volunteer{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Title:          "Helping Cleaning Bandung from Trash",
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Bandung Lover",
// 				},
// 				ContentActivity:      "We will clean the trash in Bandung",
// 				Location:             "Bandung",
// 				LinkWA:               "https://wa.me/628123456789",
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().AddDate(0, 0, 7),
// 				TargetVolunteer:      10,
// 				Status:               "active",
// 				RegisteredVolunteer:  10,
// 				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
// 				ImageURL:             "https://imageurl.com",
// 			},
// 			mockApplication: entities.Application{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				VacancyID: 1,
// 				UserID:    1,
// 			},
// 			mockError: errors.New("Unexpected Error"),
// 			expectedVolunteer: entities.Volunteer{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Title:          "Helping Cleaning Bandung from Trash",
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Bandung Lover",
// 				},
// 				ContentActivity:      "We will clean the trash in Bandung",
// 				Location:             "Bandung",
// 				LinkWA:               "https://wa.me/628123456789",
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().AddDate(0, 0, 7),
// 				TargetVolunteer:      10,
// 				Status:               "active",
// 				RegisteredVolunteer:  10,
// 				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
// 				ImageURL:             "https://imageurl.com",
// 			},
// 			expectedError: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		repository := repositories.NewVolunteerRepository(t)
// 		repository.On("FindByID", tc.mockVolunteerID).Return(tc.mockVolunteer, nil)
// 		repository.On("FindApplicationByVolunteerAndCustomer", tc.mockVolunteerID, tc.mockCustomerID).Return(tc.mockApplication, nil)
// 		repository.On("Update", tc.mockVolunteer).Return(tc.expectedVolunteer, tc.mockError)

// 		s := service.NewVolunteerService(repository)

// 		volunteer, err := s.ApplyForVolunteer(tc.mockVolunteerID, tc.mockCustomerID)

// 		assert.Equal(t, tc.expectedVolunteer, volunteer)
// 		assert.Equal(t, tc.expectedError, err)
// 	}
// }

// func TestFindApplicationByVolunteerAndCustomer(t *testing.T) {
// 	testCases := []struct {
// 		title               string
// 		mockVolunteerID     uint
// 		mockCustomerID      uint
// 		mockApplication     entities.Application
// 		mockError           error
// 		expectedApplication entities.Application
// 		expectedError       error
// 	}{
// 		{
// 			title:           "Success Find Application By Volunteer And Customer",
// 			mockVolunteerID: 1,
// 			mockCustomerID:  1,
// 			mockApplication: entities.Application{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				VacancyID: 1,
// 				UserID:    1,
// 			},
// 			mockError: nil,
// 			expectedApplication: entities.Application{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				VacancyID: 1,
// 				UserID:    1,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:               "Failed Find Application By Volunteer And Customer",
// 			mockVolunteerID:     1,
// 			mockCustomerID:      1,
// 			mockApplication:     entities.Application{},
// 			mockError:           errors.New("Unexpected Error"),
// 			expectedApplication: entities.Application{},
// 			expectedError:       errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		repository := repositories.NewVolunteerRepository(t)
// 		repository.On("FindApplicationByVolunteerAndCustomer", tc.mockVolunteerID, tc.mockCustomerID).Return(tc.mockApplication, tc.mockError)

// 		s := service.NewVolunteerService(repository)

// 		application, err := s.FindApplicationByVolunteerAndCustomer(tc.mockVolunteerID, tc.mockCustomerID)

// 		assert.Equal(t, tc.expectedApplication, application)
// 		assert.Equal(t, tc.expectedError, err)
// 	}
// }

func TestUpdateVolunteerByID(t *testing.T) {
	testCases := []struct {
		title             string
		mockID            uint
		mockVolunteer     entities.Volunteer
		mockError         error
		expectedVolunteer entities.Volunteer
		expectedError     error
	}{
		{
			title:  "Success Update Volunteer By ID",
			mockID: 1,
			mockVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			mockError: nil,
			expectedVolunteer: entities.Volunteer{
				Model: gorm.Model{
					ID: 1,
				},
				Title:          "Helping Cleaning Bandung from Trash",
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Bandung Lover",
				},
				ContentActivity:      "We will clean the trash in Bandung",
				Location:             "Bandung",
				LinkWA:               "https://wa.me/628123456789",
				StartDate:            time.Now(),
				EndDate:              time.Now().AddDate(0, 0, 7),
				TargetVolunteer:      10,
				Status:               "active",
				RegisteredVolunteer:  0,
				RegistrationDeadline: time.Now().AddDate(0, 0, 5),
				ImageURL:             "https://imageurl.com",
			},
			expectedError: nil,
		},
		{
			title:             "Failed Update Volunteer By ID",
			mockID:            1,
			mockVolunteer:     entities.Volunteer{},
			mockError:         errors.New("Unexpected Error"),
			expectedVolunteer: entities.Volunteer{},
			expectedError:     errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewVolunteerRepository(t)
		repository.On("UpdateByID", tc.mockID, tc.mockVolunteer).Return(tc.expectedVolunteer, tc.mockError)

		s := service.NewVolunteerService(repository)

		volunteer, err := s.UpdateVolunteerByID(tc.mockID, tc.mockVolunteer)

		assert.Equal(t, tc.expectedVolunteer, volunteer)
		assert.Equal(t, tc.expectedError, err)
	}
}
