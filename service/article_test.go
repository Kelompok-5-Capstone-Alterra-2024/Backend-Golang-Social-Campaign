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

func TestGetTopArticle(t *testing.T) {

	testCases := []struct {
		title        string
		mockArticles []entities.Article
		mockError    error

		expectedArticles []entities.Article
		expectedError    error
	}{
		{
			title: "Get Top Article Success",
			mockArticles: []entities.Article{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Title:    "Artikel 1",
					Content:  "Konten 1",
					ImageURL: "Image Url 1",
				},
			},
			mockError: nil,
			expectedArticles: []entities.Article{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Title:    "Artikel 1",
					Content:  "Konten 1",
					ImageURL: "Image Url 1",
				},
			},
			expectedError: nil,
		},
		{
			title:            "Get Top Article Error",
			mockArticles:     []entities.Article{},
			mockError:        errors.New("Unexpected Error"),
			expectedArticles: []entities.Article{},
			expectedError:    errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewArticleRepository(t)
		repository.On("FindTop").Return(testCase.mockArticles, testCase.mockError)

		service := service.NewArticleService(repository)
		article, err := service.GetTopArticles()

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedArticles, article)
	}
}

func TestGetByIdArticle(t *testing.T) {

	testCases := []struct {
		title                string
		mockArticlebyID      entities.Article
		mockErrorbyID        error
		id                   uint
		expectedArticlesbyID entities.Article
		expectedErrorbyID    error
	}{
		{
			title: "Get By Id Article Success",
			mockArticlebyID: entities.Article{
				Model: gorm.Model{
					ID: uint(1),
				},
				Title:    "Artikel 1",
				Content:  "Konten 1",
				ImageURL: "Image Url 1",
			},
			mockErrorbyID: nil,
			id:            uint(1),
			expectedArticlesbyID: entities.Article{
				Model: gorm.Model{
					ID: uint(1),
				},
				Title:    "Artikel 1",
				Content:  "Konten 1",
				ImageURL: "Image Url 1",
			},
			expectedErrorbyID: nil,
		},
		{
			title:                "Get By Id Article Error",
			mockArticlebyID:      entities.Article{},
			mockErrorbyID:        errors.New("Unexpected Error"),
			id:                   uint(2),
			expectedArticlesbyID: entities.Article{},
			expectedErrorbyID:    errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewArticleRepository(t)

		// kalau function findbyid dengan input testcase id ini, maka akan return mock article dan mock error
		// input ini sama dengan findbyid
		repository.On("FindByID", testCase.id).Return(testCase.mockArticlebyID, testCase.mockErrorbyID)

		service := service.NewArticleService(repository)
		article, err := service.FindByID(testCase.id)

		assert.Equal(t, testCase.expectedErrorbyID, err)
		assert.Equal(t, testCase.expectedArticlesbyID, article)
	}
}

func TestUpdateArticle(t *testing.T) {
	testCases := []struct {
		title string

		// mock variables
		mockFindArticle    entities.Article
		mockFindArticleErr error
		mockGormDb         *gorm.DB

		// function arguments
		newArticle entities.Article

		// function return
		expectedArticles entities.Article
		expectedError    error
	}{
		{
			title:              "Update Article Not Found",
			mockFindArticle:    entities.Article{},
			mockFindArticleErr: errors.New("Ga Ketemu Bang"),
			newArticle: entities.Article{
				Model: gorm.Model{
					ID: 1,
				},
				Title: "Konten Baru Bang",
			},
			expectedArticles: entities.Article{},
			expectedError:    errors.New("Ga Ketemu Bang"),
		},
		{
			title:              "Successfully Find Article But Faill Updating Article Data",
			mockFindArticle:    entities.Article{},
			mockFindArticleErr: nil,
			mockGormDb: &gorm.DB{
				Error: errors.New("Error Bang Pas Insert"),
			},
			newArticle: entities.Article{
				Model: gorm.Model{
					ID: 2,
				},
				Title: "Konten Baru Bang",
			},
			expectedArticles: entities.Article{},
			expectedError:    errors.New("Error Bang Pas Insert"),
		},
		{
			title: "Successfully Find Article And Successfully Updating Article Data",
			mockFindArticle: entities.Article{
				Model: gorm.Model{
					ID: 3,
				},
				Title: "Konten Lama Bang",
			},
			mockFindArticleErr: nil,
			mockGormDb: &gorm.DB{
				Error: nil,
			},
			newArticle: entities.Article{
				Model: gorm.Model{
					ID: 3,
				},
				Title: "Konten Baru Bang",
			},
			expectedArticles: entities.Article{
				Model: gorm.Model{
					ID: 3,
				},
				Title: "Konten Baru Bang",
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewArticleRepository(t)

		repository.On("FindByID", testCase.newArticle.ID).Return(testCase.mockFindArticle, testCase.mockFindArticleErr)

		if testCase.mockFindArticleErr == nil {
			// update content
			updatedContent := testCase.mockFindArticle
			updatedContent.Title = testCase.newArticle.Title
			updatedContent.Content = testCase.newArticle.Content
			updatedContent.ImageURL = testCase.newArticle.ImageURL

			repository.On("Save", &updatedContent).Return(testCase.mockGormDb)
		}

		service := service.NewArticleService(repository)
		article, err := service.UpdateArticle(testCase.newArticle)

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedArticles, article)
	}
}

func TestDeleteArticle(t *testing.T) {
	testCases := []struct {
		title           string
		mockDeleteError error
		id              uint
		expectedError   error
	}{
		{
			title:           "Successfully Delete Article",
			mockDeleteError: nil,
			id:              uint(1),
			expectedError:   nil,
		},
		{
			title:           "Error Delete Article",
			mockDeleteError: errors.New("Unexpected Error"),
			id:              uint(2),
			expectedError:   errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewArticleRepository(t)
		repository.On("Delete", testCase.id).Return(testCase.mockDeleteError)

		service := service.NewArticleService(repository)
		err := service.DeleteArticle(testCase.id)

		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestCreateArticle(t *testing.T) {
	testCases := []struct {
		title string

		// mock variables
		mockGormDbError error

		// function arguments
		newArticle entities.Article

		// function return
		expectedArticles entities.Article
		expectedError    error
	}{
		{
			title:           "Successfully Create Article",
			mockGormDbError: nil,
			newArticle: entities.Article{
				Model: gorm.Model{
					ID: 1,
				},
				Title: "Konten Baru Bang",
			},
			expectedArticles: entities.Article{
				Model: gorm.Model{
					ID: 1,
				},
				Title: "Konten Baru Bang",
			},
			expectedError: nil,
		},
		{
			title:           "Error Create Article",
			mockGormDbError: errors.New("Error Bang Pas Insert"),
			newArticle: entities.Article{
				Model: gorm.Model{
					ID: 2,
				},
				Title: "Konten Baru Bang",
			},
			expectedArticles: entities.Article{},
			expectedError:    errors.New("Error Bang Pas Insert"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewArticleRepository(t)

		// Configure the mock repository
		if testCase.mockGormDbError != nil {
			repository.On("Create", testCase.newArticle).Return(entities.Article{}, testCase.mockGormDbError)
		} else {
			repository.On("Create", testCase.newArticle).Return(testCase.newArticle, nil)
		}

		service := service.NewArticleService(repository)
		article, err := service.CreateArticle(testCase.newArticle)

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedArticles, article)
	}
}

func TestFindAllArticle(t *testing.T) {
	testCases := []struct {
		title string

		// mock variables
		mockArticles      []entities.Article
		mockTotalArticles int
		mockError         error

		// function arguments
		page  int
		limit int

		// function return
		expectedArticles []entities.Article
		expectedTotal    int
		expectedError    error
	}{
		{
			title: "Successfully Find All Article",
			mockArticles: []entities.Article{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Title:    "Artikel 1",
					Content:  "Konten 1",
					ImageURL: "Image Url 1",
				},
			},
			mockTotalArticles: 1,
			mockError:         nil,
			page:              1,
			limit:             1,
			expectedArticles: []entities.Article{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Title:    "Artikel 1",
					Content:  "Konten 1",
					ImageURL: "Image Url 1",
				},
			},
			expectedTotal: 1,
			expectedError: nil,
		},
		{
			title:             "Error Find All Article",
			mockArticles:      []entities.Article{},
			mockTotalArticles: 0,
			mockError:         errors.New("Unexpected Error"),
			page:              1,
			limit:             1,
			expectedArticles:  []entities.Article{},
			expectedTotal:     0,
			expectedError:     errors.New("Unexpected Error"),
		},
	}

	for _, testCase := range testCases {
		repository := repositories.NewArticleRepository(t)
		repository.On("FindAll", testCase.page, testCase.limit).Return(testCase.mockArticles, testCase.mockTotalArticles, testCase.mockError)

		service := service.NewArticleService(repository)
		articles, total, err := service.FindAll(testCase.page, testCase.limit)

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedTotal, total)
		assert.Equal(t, testCase.expectedArticles, articles)
	}
}
