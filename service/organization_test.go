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

func TestCreateOrganization(t *testing.T) {

	testCases := []struct {
		title                     string
		mockOrganization          entities.Organization
		mockOrganizationError     error
		expectedOrganization      entities.Organization
		expectedOrganizationError error
	}{
		{
			title: "Success Create Organization",
			mockOrganization: entities.Organization{
				Model: gorm.Model{
					ID: 1,
				},
				Name:        "Organization 1",
				Description: "Description 1",
				Avatar:      "Avatar 1",
				IsVerified:  true,
				StartDate:   time.Now(),
				Contact:     "Contact 1",
				Website:     "Website 1",
				Instagram:   "Instagram 1",
				NoRekening:  "No Rekening 1",
			},
			mockOrganizationError: nil,
			expectedOrganization: entities.Organization{
				Model: gorm.Model{
					ID: 1,
				},
				Name:        "Organization 1",
				Description: "Description 1",
				Avatar:      "Avatar 1",
				IsVerified:  true,
				StartDate:   time.Now(),
				Contact:     "Contact 1",
				Website:     "Website 1",
				Instagram:   "Instagram 1",
				NoRekening:  "No Rekening 1",
			},
			expectedOrganizationError: nil,
		},
		{
			title:                     "Failed Create Organization",
			mockOrganization:          entities.Organization{},
			mockOrganizationError:     errors.New("Unexpected Error"),
			expectedOrganization:      entities.Organization{},
			expectedOrganizationError: errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewOrganizationRepository(t)
		repository.On("Save", tc.mockOrganization).Return(tc.mockOrganization, tc.mockOrganizationError)

		s := service.NewOrganizationService(repository)
		organization, err := s.CreateOrganization(tc.mockOrganization)

		assert.Equal(t, tc.expectedOrganization, organization)
		assert.Equal(t, tc.expectedOrganizationError, err)
	}
}

// func TestFindAllOrganization(t *testing.T) {

// 	testCases := []struct {
// 		title                      string
// 		mockOrganizations          []entities.Organization
// 		mockOrganizationsError     error
// 		expectedOrganizations      []entities.Organization
// 		expectedOrganizationsError error
// 	}{
// 		{
// 			title: "Success Find All Organizations",
// 			mockOrganizations: []entities.Organization{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name:        "Organization 1",
// 					Description: "Description 1",
// 					Avatar:      "Avatar 1",
// 					IsVerified:  true,
// 					StartDate:   time.Now(),
// 					Contact:     "Contact 1",
// 					Website:     "Website 1",
// 					Instagram:   "Instagram 1",
// 					NoRekening:  "No Rekening 1",
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					Name:        "Organization 2",
// 					Description: "Description 2",
// 					Avatar:      "Avatar 2",
// 					IsVerified:  true,
// 					StartDate:   time.Now(),
// 					Contact:     "Contact 2",
// 					Website:     "Website 2",
// 					Instagram:   "Instagram 2",
// 					NoRekening:  "No Rekening 2",
// 				},
// 			},
// 			mockOrganizationsError: nil,
// 			expectedOrganizations: []entities.Organization{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name:        "Organization 1",
// 					Description: "Description 1",
// 					Avatar:      "Avatar 1",
// 					IsVerified:  true,
// 					StartDate:   time.Now(),
// 					Contact:     "Contact 1",
// 					Website:     "Website 1",
// 					Instagram:   "Instagram 1",
// 					NoRekening:  "No Rekening 1",
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					Name:        "Organization 2",
// 					Description: "Description 2",
// 					Avatar:      "Avatar 2",
// 					IsVerified:  true,
// 					StartDate:   time.Now(),
// 					Contact:     "Contact 2",
// 					Website:     "Website 2",
// 					Instagram:   "Instagram 2",
// 					NoRekening:  "No Rekening 2",
// 				},
// 			},
// 			expectedOrganizationsError: nil,
// 		},
// 		{
// 			title:                      "Failed Find All Organizations",
// 			mockOrganizations:          []entities.Organization{},
// 			mockOrganizationsError:     errors.New("Unexpected Error"),
// 			expectedOrganizations:      []entities.Organization{},
// 			expectedOrganizationsError: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		repository := repositories.NewOrganizationRepository(t)
// 		repository.On("FindAll").Return(tc.mockOrganizations, tc.mockOrganizationsError)

// 		s := service.NewOrganizationService(repository)
// 		organizations, err := s.FindAllOrganization()

// 		assert.Equal(t, tc.expectedOrganizations, organizations)
// 		assert.Equal(t, tc.expectedOrganizationsError, err)
// 	}
// }

func TestFindOrganizationByID(t *testing.T) {

	testCases := []struct {
		title                   string
		mockOrganizationID      int
		mockOrganization        entities.Organization
		mockOrganizationError   error
		expectedOrganization    entities.Organization
		expectedOrganizationErr error
	}{
		{
			title:              "Success Find Organization By ID",
			mockOrganizationID: 1,
			mockOrganization: entities.Organization{
				Model: gorm.Model{
					ID: 1,
				},
				Name:        "Organization 1",
				Description: "Description 1",
				Avatar:      "Avatar 1",
				IsVerified:  true,
				StartDate:   time.Now(),
				Contact:     "Contact 1",
				Website:     "Website 1",
				Instagram:   "Instagram 1",
				NoRekening:  "No Rekening 1",
			},
			mockOrganizationError: nil,
			expectedOrganization: entities.Organization{
				Model: gorm.Model{
					ID: 1,
				},
				Name:        "Organization 1",
				Description: "Description 1",
				Avatar:      "Avatar 1",
				IsVerified:  true,
				StartDate:   time.Now(),
				Contact:     "Contact 1",
				Website:     "Website 1",
				Instagram:   "Instagram 1",
				NoRekening:  "No Rekening 1",
			},
			expectedOrganizationErr: nil,
		},
		{
			title:                   "Failed Find Organization By ID",
			mockOrganizationID:      1,
			mockOrganization:        entities.Organization{},
			mockOrganizationError:   errors.New("Unexpected Error"),
			expectedOrganization:    entities.Organization{},
			expectedOrganizationErr: errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		repository := repositories.NewOrganizationRepository(t)
		repository.On("FindByID", tc.mockOrganizationID).Return(tc.mockOrganization, tc.mockOrganizationError)

		s := service.NewOrganizationService(repository)
		organization, err := s.FindOrganizationByID(tc.mockOrganizationID)

		assert.Equal(t, tc.expectedOrganization, organization)
		assert.Equal(t, tc.expectedOrganizationErr, err)
	}
}

// func TestFindFundraisingByOrganizationID(t *testing.T) {

// 	testCases := []struct {
// 		title                          string
// 		mockOrganizationID             int
// 		mockFundraisings               []entities.Fundraising
// 		mockFundraisingsError          error
// 		expectedFundraisings           []entities.Fundraising
// 		expectedFundraisingsError      error
// 		expectedFundraisingsErrorCount int
// 	}{
// 		{
// 			title:              "Success Find Fundraising By Organization ID",
// 			mockOrganizationID: 1,
// 			mockFundraisings: []entities.Fundraising{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Title:       "Fundraising 1",
// 					Description: "Description 1",
// 					Target:      1000000,
// 					Current:     500000,
// 					StartDate:   time.Now(),
// 					EndDate:     time.Now(),
// 					IsVerified:  true,
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					Title:       "Fundraising 2",
// 					Description: "Description 2",
// 					Target:      2000000,
// 					Current:     1000000,
// 					StartDate:   time.Now(),
// 					EndDate:     time.Now(),
// 					IsVerified:  true,
// 				},
// 			},
// 			mockFundraisingsError: nil,
// 			expectedFundraisings: []entities.Fundraising{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Title:       "Fundraising 1",
// 					Description: "Description 1",
// 					Target:      1000000,
// 					Current:     500000,
// 					StartDate:   time.Now(),
// 					EndDate:     time.Now(),
// 					IsVerified:  true,
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					Title:       "Fundraising 2",
// 					Description: "Description 2",
// 					Target:      2000000,
// 					Current:     1000000,
// 					StartDate:   time.Now(),
// 					EndDate:     time.Now(),
// 					IsVerified:  true,
// 				},
// 			},
// 			expectedFundraisingsError:      nil,
// 			expectedFundraisingsErrorCount: 0,
// 		},
// 		{
// 			title:                          "Failed Find Fundraising By Organization ID",
// 			mockOrganizationID:             1,
// 			mockFundraisings:               []entities.Fundraising{},
// 			mockFundraisingsError:          errors.New("Unexpected Error"),
// 			expectedFundraisings:           []entities.Fundraising{},
// 			expectedFundraisingsError:      errors.New("Unexpected Error"),
// 			expectedFundraisingsErrorCount: 1,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		repository := repositories.NewOrganizationRepository(t)
// 		repository.On("FindFundraisingByOrganizationID", tc.mockOrganizationID, 0, 0).Return(tc.mockFundraisings, tc.mockFundraisingsError)

// 		s := service.NewOrganizationService(repository)
// 		fundraisings, err := s.FindFundraisingByOrganizationID(tc.mockOrganizationID, 0, 0)

// 		assert.Equal(t, tc.expectedFundraisings, fundraings)
// 		assert.Equal(t, tc.expectedFundraisingsError, err)
// 	}
// }

// func TestUpdateOrganization(t *testing.T) {

// 	testCases := []struct {
// 		title                     string
// 		mockOrganization          entities.Organization
// 		mockOrganizationError     error
// 		expectedOrganization      entities.Organization
// 		expectedOrganizationError error
// 	}{
// 		{
// 			title: "Success Update Organization",
// 			mockOrganization: entities.Organization{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Name:        "Organization 1",
// 				Description: "Description 1",
// 				Avatar:      "Avatar 1",
// 				IsVerified:  true,
// 				StartDate:   time.Now(),
// 				Contact:     "Contact 1",
// 				Website:     "Website 1",
// 				Instagram:   "Instagram 1",
// 				NoRekening:  "No Rekening 1",
// 			},
// 			mockOrganizationError: nil,
// 			expectedOrganization: entities.Organization{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Name:        "Organization 1",
// 				Description: "Description 1",
// 				Avatar:      "Avatar 1",
// 				IsVerified:  true,
// 				StartDate:   time.Now(),
// 				Contact:     "Contact 1",
// 				Website:     "Website 1",
// 				Instagram:   "Instagram 1",
// 				NoRekening:  "No Rekening 1",
// 			},
// 			expectedOrganizationError: nil,
// 		},
// 		{
// 			title:                     "Failed Update Organization",
// 			mockOrganization:          entities.Organization{},
// 			mockOrganizationError:     errors.New("Unexpected Error"),
// 			expectedOrganization:      entities.Organization{},
// 			expectedOrganizationError: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		repository := repositories.NewOrganizationRepository(t)
// 		repository.On("Save", tc.mockOrganization).Return(tc.mockOrganization, tc.mockOrganizationError)

// 		s := service.NewOrganizationService(repository)
// 		organization, err := s.UpdateOrganization(tc.mockOrganization)

// 		assert.Equal(t, tc.expectedOrganization, organization)
// 		assert.Equal(t, tc.expectedOrganizationError, err)
// 	}
// }
