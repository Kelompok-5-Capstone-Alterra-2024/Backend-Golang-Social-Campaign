package service_test

import (
	"capstone/entities"
	"capstone/mocks/repositories"
	"capstone/service"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// func TestCreateFundraising(t *testing.T) {

// 	testCases := []struct {
// 		title                    string
// 		mockFundraising          entities.Fundraising
// 		mockErrorFundraising     error
// 		expectedFundraising      entities.Fundraising
// 		expectedErrorFundraising error
// 	}{
// 		{
// 			title: "Success Create Fundraising",
// 			mockFundraising: entities.Fundraising{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				FundraisingCategoryID: 1,
// 				FundraisingCategory: entities.FundraisingCategory{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Category 1",
// 				},
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Organization 1",
// 				},
// 				Title:           "Fundraising 1",
// 				ImageUrl:        "image_url",
// 				Description:     "Description",
// 				Status:          "status",
// 				GoalAmount:      1000,
// 				CurrentProgress: 0,
// 				StartDate:       &time.Time{},
// 				EndDate:         &time.Time{},
// 			},
// 			mockErrorFundraising: nil,
// 			expectedFundraising: entities.Fundraising{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				FundraisingCategoryID: 1,
// 				FundraisingCategory: entities.FundraisingCategory{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Category 1",
// 				},
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Organization 1",
// 				},
// 				Title:           "Fundraising 1",
// 				ImageUrl:        "image_url",
// 				Description:     "Description",
// 				Status:          "status",
// 				GoalAmount:      1000,
// 				CurrentProgress: 0,
// 				StartDate:       &time.Time{},
// 				EndDate:         &time.Time{},
// 			},
// 			expectedErrorFundraising: nil,
// 		},
// 		{
// 			title:                    "Failed Create Fundraising",
// 			mockFundraising:          entities.Fundraising{},
// 			mockErrorFundraising:     errors.New("Unexpected Error"),
// 			expectedFundraising:      entities.Fundraising{},
// 			expectedErrorFundraising: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewFundraisingRepository(t)
// 		repository.On("Create", mock.Anything).Return(testCase.mockFundraising, testCase.mockErrorFundraising)

// 		service := service.NewFundraisingService(repository)
// 		fundraising, err := service.CreateFundraising(testCase.mockFundraising)

// 		assert.Equal(t, testCase.expectedErrorFundraising, err)
// 		assert.Equal(t, testCase.expectedFundraising, fundraising)
// 	}
// }

// func TestUpdateFundraising(t *testing.T) {

// 	testCases := []struct {
// 		title                    string
// 		mockFundraising          entities.Fundraising
// 		mockErrorFundraising     error
// 		expectedFundraising      entities.Fundraising
// 		expectedErrorFundraising error
// 	}{
// 		{
// 			title: "Success Update Fundraising",
// 			mockFundraising: entities.Fundraising{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				FundraisingCategoryID: 1,
// 				FundraisingCategory: entities.FundraisingCategory{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Category 1",
// 				},
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Organization 1",
// 				},
// 				Title:           "Fundraising 1",
// 				ImageUrl:        "image_url",
// 				Description:     "Description",
// 				Status:          "status",
// 				GoalAmount:      1000,
// 				CurrentProgress: 0,
// 				StartDate:       &time.Time{},
// 				EndDate:         &time.Time{},
// 			},
// 			mockErrorFundraising: nil,
// 			expectedFundraising: entities.Fundraising{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				FundraisingCategoryID: 1,
// 				FundraisingCategory: entities.FundraisingCategory{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Category 1",
// 				},
// 				OrganizationID: 1,
// 				Organization: entities.Organization{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Name: "Organization 1",
// 				},
// 				Title:           "Fundraising 1",
// 				ImageUrl:        "image_url",
// 				Description:     "Description",
// 				Status:          "status",
// 				GoalAmount:      1000,
// 				CurrentProgress: 0,
// 				StartDate:       &time.Time{},
// 				EndDate:         &time.Time{},
// 			},
// 			expectedErrorFundraising: nil,
// 		},
// 		{
// 			title:                    "Failed Update Fundraising",
// 			mockFundraising:          entities.Fundraising{},
// 			mockErrorFundraising:     errors.New("Unexpected Error"),
// 			expectedFundraising:      entities.Fundraising{},
// 			expectedErrorFundraising: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewFundraisingRepository(t)
// 		repository.On("Update", mock.Anything).Return(testCase.mockFundraising, testCase.mockErrorFundraising)

// 		service := service.NewFundraisingService(repository)
// 		fundraising, err := service.UpdateFundraising(testCase.mockFundraising)

// 		assert.Equal(t, testCase.expectedErrorFundraising, err)
// 		assert.Equal(t, testCase.expectedFundraising, fundraising)
// 	}
// }

func TestFindFundraisingByID(t *testing.T) {

	testCases := []struct {
		title                    string
		mockFundraising          entities.Fundraising
		mockErrorFundraising     error
		expectedFundraising      entities.Fundraising
		expectedErrorFundraising error
	}{
		{
			title: "Success Find Fundraising By ID",
			mockFundraising: entities.Fundraising{
				Model: gorm.Model{
					ID: 1,
				},
				FundraisingCategoryID: 1,
				FundraisingCategory: entities.FundraisingCategory{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Category 1",
				},
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Organization 1",
				},
				Title:           "Fundraising 1",
				ImageUrl:        "image_url",
				Description:     "Description",
				Status:          "status",
				GoalAmount:      1000,
				CurrentProgress: 0,
				StartDate:       &time.Time{},
				EndDate:         &time.Time{},
			},
			mockErrorFundraising: nil,
			expectedFundraising: entities.Fundraising{
				Model: gorm.Model{
					ID: 1,
				},
				FundraisingCategoryID: 1,
				FundraisingCategory: entities.FundraisingCategory{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Category 1",
				},
				OrganizationID: 1,
				Organization: entities.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Organization 1",
				},
				Title:           "Fundraising 1",
				ImageUrl:        "image_url",
				Description:     "Description",
				Status:          "status",
				GoalAmount:      1000,
				CurrentProgress: 0,
				StartDate:       &time.Time{},
				EndDate:         &time.Time{},
			},
			expectedErrorFundraising: nil,
		},
		{
			title:                    "Failed Find Fundraising By ID",
			mockFundraising:          entities.Fundraising{},
			mockErrorFundraising:     errors.New("Unexpected Error"),
			expectedFundraising:      entities.Fundraising{},
			expectedErrorFundraising: errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewFundraisingRepository(t)
		repository.On("FindByID", mock.Anything).Return(testCase.mockFundraising, testCase.mockErrorFundraising)

		service := service.NewFundraisingService(repository)
		fundraising, err := service.FindFundraisingByID(1)

		assert.Equal(t, testCase.expectedErrorFundraising, err)
		assert.Equal(t, testCase.expectedFundraising, fundraising)
	}
}

// func TestFindFundraisings(t *testing.T) {

// 	testCases := []struct {
// 		title                    string
// 		mockFundraising          []entities.Fundraising
// 		mockErrorFundraising     error
// 		expectedFundraising      []entities.Fundraising
// 		expectedErrorFundraising error
// 	}{
// 		{
// 			title: "Success Find Fundraisings",
// 			mockFundraising: []entities.Fundraising{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					FundraisingCategoryID: 1,
// 					FundraisingCategory: entities.FundraisingCategory{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Category 1",
// 					},
// 					OrganizationID: 1,
// 					Organization: entities.Organization{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Organization 1",
// 					},
// 					Title:           "Fundraising 1",
// 					ImageUrl:        "image_url",
// 					Description:     "Description",
// 					Status:          "status",
// 					GoalAmount:      1000,
// 					CurrentProgress: 0,
// 					StartDate:       &time.Time{},
// 					EndDate:         &time.Time{},
// 				},
// 			},
// 			mockErrorFundraising: nil,
// 			expectedFundraising: []entities.Fundraising{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					FundraisingCategoryID: 1,
// 					FundraisingCategory: entities.FundraisingCategory{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Category 1",
// 					},
// 					OrganizationID: 1,
// 					Organization: entities.Organization{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Organization 1",
// 					},
// 					Title:           "Fundraising 1",
// 					ImageUrl:        "image_url",
// 					Description:     "Description",
// 					Status:          "status",
// 					GoalAmount:      1000,
// 					CurrentProgress: 0,
// 					StartDate:       &time.Time{},
// 					EndDate:         &time.Time{},
// 				},
// 			},
// 			expectedErrorFundraising: nil,
// 		},
// 		{
// 			title:                    "Failed Find Fundraisings",
// 			mockFundraising:          []entities.Fundraising{},
// 			mockErrorFundraising:     errors.New("Unexpected Error"),
// 			expectedFundraising:      []entities.Fundraising{},
// 			expectedErrorFundraising: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewFundraisingRepository(t)
// 		repository.On("FindAll", mock.Anything, mock.Anything).Return(testCase.mockFundraising, testCase.mockErrorFundraising)

// 		service := service.NewFundraisingService(repository)
// 		fundraising, err := service.FindFundraisings(1, 1)

// 		assert.Equal(t, testCase.expectedErrorFundraising, err)
// 		assert.Equal(t, testCase.expectedFundraising, fundraising)
// 	}
// }

func TestFindTopFundraisings(t *testing.T) {

	testCases := []struct {
		title                    string
		mockFundraising          []entities.Fundraising
		mockErrorFundraising     error
		expectedFundraising      []entities.Fundraising
		expectedErrorFundraising error
	}{
		{
			title: "Success Find Top Fundraisings",
			mockFundraising: []entities.Fundraising{
				{
					Model: gorm.Model{
						ID: 1,
					},
					FundraisingCategoryID: 1,
					FundraisingCategory: entities.FundraisingCategory{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Category 1",
					},
					OrganizationID: 1,
					Organization: entities.Organization{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Organization 1",
					},
					Title:           "Fundraising 1",
					ImageUrl:        "image_url",
					Description:     "Description",
					Status:          "status",
					GoalAmount:      1000,
					CurrentProgress: 0,
					StartDate:       &time.Time{},
					EndDate:         &time.Time{},
				},
			},
			mockErrorFundraising: nil,
			expectedFundraising: []entities.Fundraising{
				{
					Model: gorm.Model{
						ID: 1,
					},
					FundraisingCategoryID: 1,
					FundraisingCategory: entities.FundraisingCategory{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Category 1",
					},
					OrganizationID: 1,
					Organization: entities.Organization{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "Organization 1",
					},
					Title:           "Fundraising 1",
					ImageUrl:        "image_url",
					Description:     "Description",
					Status:          "status",
					GoalAmount:      1000,
					CurrentProgress: 0,
					StartDate:       &time.Time{},
					EndDate:         &time.Time{},
				},
			},
			expectedErrorFundraising: nil,
		},
		{
			title:                    "Failed Find Top Fundraisings",
			mockFundraising:          []entities.Fundraising{},
			mockErrorFundraising:     errors.New("Unexpected Error"),
			expectedFundraising:      []entities.Fundraising{},
			expectedErrorFundraising: errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewFundraisingRepository(t)
		repository.On("FindTopFundraisings").Return(testCase.mockFundraising, testCase.mockErrorFundraising)

		service := service.NewFundraisingService(repository)
		fundraising, err := service.FindTopFundraisings()

		assert.Equal(t, testCase.expectedErrorFundraising, err)
		assert.Equal(t, testCase.expectedFundraising, fundraising)
	}
}

func TestFindAllFundraisingCategories(t *testing.T) {

	testCases := []struct {
		title                    string
		mockFundraisingCategory  []entities.FundraisingCategory
		mockErrorFundraising     error
		expectedFundraising      []entities.FundraisingCategory
		expectedErrorFundraising error
	}{
		{
			title: "Success Find All Fundraising Categories",
			mockFundraisingCategory: []entities.FundraisingCategory{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Category 1",
				},
			},
			mockErrorFundraising: nil,
			expectedFundraising: []entities.FundraisingCategory{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "Category 1",
				},
			},
			expectedErrorFundraising: nil,
		},
		{
			title:                    "Failed Find All Fundraising Categories",
			mockFundraisingCategory:  []entities.FundraisingCategory{},
			mockErrorFundraising:     errors.New("Unexpected Error"),
			expectedFundraising:      []entities.FundraisingCategory{},
			expectedErrorFundraising: errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewFundraisingRepository(t)
		repository.On("FindAllCategories").Return(testCase.mockFundraisingCategory, testCase.mockErrorFundraising)

		service := service.NewFundraisingService(repository)
		fundraisingCategory, err := service.FindAllFundraisingCategories()

		assert.Equal(t, testCase.expectedErrorFundraising, err)
		assert.Equal(t, testCase.expectedFundraising, fundraisingCategory)
	}
}

// func TestFindFundraisingByCategoryID(t *testing.T) {

// 	testCases := []struct {
// 		title                    string
// 		mockFundraising          []entities.Fundraising
// 		mockErrorFundraising     error
// 		expectedFundraising      []entities.Fundraising
// 		expectedErrorFundraising error
// 	}{
// 		{
// 			title: "Success Find Fundraising By Category ID",
// 			mockFundraising: []entities.Fundraising{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					FundraisingCategoryID: 1,
// 					FundraisingCategory: entities.FundraisingCategory{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Category 1",
// 					},
// 					OrganizationID: 1,
// 					Organization: entities.Organization{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Organization 1",
// 					},
// 					Title:           "Fundraising 1",
// 					ImageUrl:        "image_url",
// 					Description:     "Description",
// 					Status:          "status",
// 					GoalAmount:      1000,
// 					CurrentProgress: 0,
// 					StartDate:       &time.Time{},
// 					EndDate:         &time.Time{},
// 				},
// 			},
// 			mockErrorFundraising: nil,
// 			expectedFundraising: []entities.Fundraising{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					FundraisingCategoryID: 1,
// 					FundraisingCategory: entities.FundraisingCategory{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Category 1",
// 					},
// 					OrganizationID: 1,
// 					Organization: entities.Organization{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Name: "Organization 1",
// 					},
// 					Title:           "Fundraising 1",
// 					ImageUrl:        "image_url",
// 					Description:     "Description",
// 					Status:          "status",
// 					GoalAmount:      1000,
// 					CurrentProgress: 0,
// 					StartDate:       &time.Time{},
// 					EndDate:         &time.Time{},
// 				},
// 			},
// 			expectedErrorFundraising: nil,
// 		},
// 		{
// 			title:                    "Failed Find Fundraising By Category ID",
// 			mockFundraising:          []entities.Fundraising{},
// 			mockErrorFundraising:     errors.New("Unexpected Error"),
// 			expectedFundraising:      []entities.Fundraising{},
// 			expectedErrorFundraising: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewFundraisingRepository(t)
// 		repository.On("FindByCategoryID", mock.Anything, mock.Anything).Return(testCase.mockFundraising, testCase.mockErrorFundraising)

// 		service := service.NewFundraisingService(repository)
// 		fundraising, err := service.FindFundraisingByCategoryID(1, 1, 1)

// 		assert.Equal(t, testCase.expectedErrorFundraising, err)
// 		assert.Equal(t, testCase.expectedFundraising, fundraising)
// 	}
// }
