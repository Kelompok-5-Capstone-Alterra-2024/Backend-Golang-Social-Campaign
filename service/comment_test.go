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

func TestCreateComment(t *testing.T) {

	testCases := []struct {
		title string

		mockGormDBError error

		newComment entities.Comment

		expectedComment entities.Comment
		expectedError   error
	}{
		{
			title:           "Successfully Create Comment",
			mockGormDBError: nil,
			newComment: entities.Comment{
				Model: gorm.Model{
					ID: 1,
				},
				ArticleID: 1,
				Comment:   "This is a comment",
			},
			expectedComment: entities.Comment{
				Model: gorm.Model{
					ID: 1,
				},
				ArticleID: 1,
				Comment:   "This is a comment",
			},
			expectedError: nil,
		},
		{
			title:           "Failed Create Comment",
			mockGormDBError: errors.New("Unexpected Error"),
			newComment: entities.Comment{
				Model: gorm.Model{
					ID: 2,
				},
				ArticleID: 2,
				Comment:   "This is a comment",
			},
			expectedComment: entities.Comment{},
			expectedError:   errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			repository := repositories.NewCommentRepository(t)

			if tc.mockGormDBError != nil {
				repository.On("Create", tc.newComment).Return(entities.Comment{}, tc.mockGormDBError)
			} else {
				repository.On("Create", tc.newComment).Return(tc.expectedComment, nil)
			}

			commentService := service.NewCommentService(repository)

			comment, err := commentService.CreateComment(tc.newComment)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}

func TestUpdateComment(t *testing.T) {

	testCases := []struct {
		title string

		mockGormDBError error

		newComment entities.Comment

		expectedComment entities.Comment
		expectedError   error
	}{
		{
			title:           "Successfully Update Comment",
			mockGormDBError: nil,
			newComment: entities.Comment{
				Model: gorm.Model{
					ID: 1,
				},
				ArticleID: 1,
				Comment:   "This is a comment",
			},
			expectedComment: entities.Comment{
				Model: gorm.Model{
					ID: 1,
				},
				ArticleID: 1,
				Comment:   "This is a comment",
			},
			expectedError: nil,
		},
		{
			title:           "Failed Update Comment",
			mockGormDBError: errors.New("Unexpected Error"),
			newComment: entities.Comment{
				Model: gorm.Model{
					ID: 2,
				},
				ArticleID: 2,
				Comment:   "This is a comment",
			},
			expectedComment: entities.Comment{},
			expectedError:   errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			repository := repositories.NewCommentRepository(t)

			if tc.mockGormDBError != nil {
				repository.On("Update", tc.newComment).Return(entities.Comment{}, tc.mockGormDBError)
			} else {
				repository.On("Update", tc.newComment).Return(tc.expectedComment, nil)
			}

			commentService := service.NewCommentService(repository)

			comment, err := commentService.UpdateComment(tc.newComment)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}

func TestFindCommentByID(t *testing.T) {

	testCases := []struct {
		title string

		mockGormDBError error

		commentID uint

		expectedComment entities.Comment
		expectedError   error
	}{
		{
			title:           "Successfully Find Comment By ID",
			mockGormDBError: nil,
			commentID:       1,
			expectedComment: entities.Comment{
				Model: gorm.Model{
					ID: 1,
				},
				ArticleID: 1,
				Comment:   "This is a comment",
			},
			expectedError: nil,
		},
		{
			title:           "Failed Find Comment By ID",
			mockGormDBError: errors.New("Unexpected Error"),
			commentID:       2,
			expectedComment: entities.Comment{},
			expectedError:   errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			repository := repositories.NewCommentRepository(t)

			if tc.mockGormDBError != nil {
				repository.On("FindByID", tc.commentID).Return(entities.Comment{}, tc.mockGormDBError)
			} else {
				repository.On("FindByID", tc.commentID).Return(tc.expectedComment, nil)
			}

			commentService := service.NewCommentService(repository)

			comment, err := commentService.FindByID(tc.commentID)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}

func TestGetAllByArticleID(t *testing.T) {

	testCases := []struct {
		title string

		mockGormDBError error

		articleID uint
		page      int
		limit     int

		expectedComments []entities.Comment
		expectedTotal    int
		expectedError    error
	}{
		{
			title:           "Successfully Get All Comments By Article ID",
			mockGormDBError: nil,
			articleID:       1,
			page:            1,
			limit:           10,
			expectedComments: []entities.Comment{
				{
					Model: gorm.Model{
						ID: 1,
					},
					ArticleID: 1,
					Comment:   "This is a comment",
				},
			},
			expectedTotal: 1,
			expectedError: nil,
		},
		{
			title:            "Failed Get All Comments By Article ID",
			mockGormDBError:  errors.New("Unexpected Error"),
			articleID:        2,
			page:             1,
			limit:            10,
			expectedComments: []entities.Comment{},
			expectedTotal:    0,
			expectedError:    errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			repository := repositories.NewCommentRepository(t)

			if tc.mockGormDBError != nil {
				repository.On("FindAllByArticleID", tc.articleID, tc.page, tc.limit).Return([]entities.Comment{}, 0, tc.mockGormDBError)
			} else {
				repository.On("FindAllByArticleID", tc.articleID, tc.page, tc.limit).Return(tc.expectedComments, tc.expectedTotal, nil)
			}

			commentService := service.NewCommentService(repository)

			comments, total, err := commentService.GetAllByArticleID(tc.articleID, tc.page, tc.limit)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComments, comments)
			assert.Equal(t, tc.expectedTotal, total)
		})
	}
}

func TestDeleteComment(t *testing.T) {

	testCases := []struct {
		title string

		mockGormDBError error

		commentID uint

		expectedError error
	}{
		{
			title:           "Successfully Delete Comment",
			mockGormDBError: nil,
			commentID:       1,
			expectedError:   nil,
		},
		{
			title:           "Failed Delete Comment",
			mockGormDBError: errors.New("Unexpected Error"),
			commentID:       2,
			expectedError:   errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			repository := repositories.NewCommentRepository(t)

			if tc.mockGormDBError != nil {
				repository.On("Delete", tc.commentID).Return(tc.mockGormDBError)
			} else {
				repository.On("Delete", tc.commentID).Return(nil)
			}

			commentService := service.NewCommentService(repository)

			err := commentService.DeleteComment(tc.commentID)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetAllComments(t *testing.T) {

	testCases := []struct {
		title string

		mockGormDBError error

		expectedComments []entities.Comment
		expectedError    error
	}{
		{
			title:           "Successfully Get All Comments",
			mockGormDBError: nil,
			expectedComments: []entities.Comment{
				{
					Model: gorm.Model{
						ID: 1,
					},
					ArticleID: 1,
					Comment:   "This is a comment",
				},
			},
			expectedError: nil,
		},
		{
			title:            "Failed Get All Comments",
			mockGormDBError:  errors.New("Unexpected Error"),
			expectedComments: []entities.Comment{},
			expectedError:    errors.New("Unexpected Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			repository := repositories.NewCommentRepository(t)

			if tc.mockGormDBError != nil {
				repository.On("FindAll").Return([]entities.Comment{}, tc.mockGormDBError)
			} else {
				repository.On("FindAll").Return(tc.expectedComments, nil)
			}

			commentService := service.NewCommentService(repository)

			comments, err := commentService.GetAllComments()
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComments, comments)
		})
	}
}
