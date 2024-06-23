package service_test

// func TestCreateLikesComment(t *testing.T) {

// 	testCases := []struct {
// 		title                     string
// 		mockLikesComment          entities.LikesComment
// 		mockErrorLikesComment     error
// 		expectedLikesComment      entities.LikesComment
// 		expectedErrorLikesComment error
// 	}{
// 		{
// 			title: "Success Create Likes Comment",
// 			mockLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			mockErrorLikesComment: nil,
// 			expectedLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			expectedErrorLikesComment: nil,
// 		},
// 		{
// 			title: "Failed Create Likes Comment",
// 			mockLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			mockErrorLikesComment:     errors.New("Failed to create likes comment"),
// 			expectedLikesComment:      entities.LikesComment{},
// 			expectedErrorLikesComment: errors.New("Failed to create likes comment"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewLikesCommentRepository(t)
// 		repository.On("Create", tt.mockLikesComment).Return(tt.mockErrorLikesComment)

// 		service := service.NewLikesCommentService(repository)
// 		likesComment, err := service.LikesComment(tt.mockLikesComment.CommentID, tt.mockLikesComment.UserID)

// 		assert.Equal(t, tt.expectedErrorLikesComment, err)
// 		assert.Equal(t, tt.expectedLikesComment, likesComment)
// 	}
// }

// func TestDeleteLikesComment(t *testing.T) {

// 	testCases := []struct {
// 		title                     string
// 		mockLikesComment          entities.LikesComment
// 		mockErrorLikesComment     error
// 		expectedLikesComment      entities.LikesComment
// 		expectedErrorLikesComment error
// 	}{
// 		{
// 			title: "Success Delete Likes Comment",
// 			mockLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			mockErrorLikesComment: nil,
// 			expectedLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			expectedErrorLikesComment: nil,
// 		},
// 		{
// 			title: "Failed Delete Likes Comment",
// 			mockLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			mockErrorLikesComment:     errors.New("Failed to delete likes comment"),
// 			expectedLikesComment:      entities.LikesComment{},
// 			expectedErrorLikesComment: errors.New("Failed to delete likes comment"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewLikesCommentRepository(t)
// 		repository.On("Delete", tt.mockLikesComment.CommentID, tt.mockLikesComment.UserID).Return(tt.mockErrorLikesComment)

// 		service := service.NewLikesCommentService(repository)
// 		likesComment, err := service.DeleteLikesComment(tt.mockLikesComment.CommentID, tt.mockLikesComment.UserID)

// 		assert.Equal(t, tt.expectedErrorLikesComment, err)
// 		assert.Equal(t, tt.expectedLikesComment, likesComment)
// 	}
// }

// func TestGetLikesCommentByID(t *testing.T) {

// 	testCases := []struct {
// 		title                     string
// 		mockLikesComment          entities.LikesComment
// 		mockErrorLikesComment     error
// 		expectedLikesComment      entities.LikesComment
// 		expectedErrorLikesComment error
// 	}{
// 		{
// 			title: "Success Get Likes Comment By ID",
// 			mockLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			mockErrorLikesComment: nil,
// 			expectedLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			expectedErrorLikesComment: nil,
// 		},
// 		{
// 			title: "Failed Get Likes Comment By ID",
// 			mockLikesComment: entities.LikesComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				CommentID: 1,
// 				UserID:    1,
// 			},
// 			mockErrorLikesComment:     errors.New("Failed to get likes comment by ID"),
// 			expectedLikesComment:      entities.LikesComment{},
// 			expectedErrorLikesComment: errors.New("Failed to get likes comment by ID"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewLikesCommentRepository(t)
// 		repository.On("FindByID", tt.mockLikesComment.ID).Return(tt.mockLikesComment, tt.mockErrorLikesComment)

// 		service := service.NewLikesCommentService(repository)
// 		likesComment, err := service.GetLikesCommentByID(tt.mockLikesComment.ID)

// 		assert.Equal(t, tt.expectedErrorLikesComment, err)
// 		assert.Equal(t, tt.expectedLikesComment, likesComment)
// 	}
// }

// func TestGetAllLikesComments(t *testing.T) {

// 	testCases := []struct {
// 		title                     string
// 		mockLikesComments         []entities.LikesComment
// 		mockErrorLikesComments    error
// 		expectedLikesComments     []entities.LikesComment
// 		expectedErrorLikesComment error
// 	}{
// 		{
// 			title: "Success Get All Likes Comments",
// 			mockLikesComments: []entities.LikesComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					CommentID: 1,
// 					UserID:    1,
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					CommentID: 2,
// 					UserID:    2,
// 				},
// 			},
// 			mockErrorLikesComments: nil,
// 			expectedLikesComments: []entities.LikesComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					CommentID: 1,
// 					UserID:    1,
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					CommentID: 2,
// 					UserID:    2,
// 				},
// 			},
// 			expectedErrorLikesComment: nil,
// 		},
// 		{
// 			title: "Failed Get All Likes Comments",
// 			mockLikesComments: []entities.LikesComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					CommentID: 1,
// 					UserID:    1,
// 				},
// 				{
// 					Model: gorm.Model{
// 						ID: 2,
// 					},
// 					CommentID: 2,
// 					UserID:    2,
// 				},
// 			},
// 			mockErrorLikesComments:    errors.New("Failed to get all likes comments"),
// 			expectedLikesComments:     []entities.LikesComment{},
// 			expectedErrorLikesComment: errors.New("Failed to get all likes comments"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewLikesCommentRepository(t)
// 		repository.On("FindAll").Return(tt.mockLikesComments, tt.mockErrorLikesComments)

// 		service := service.NewLikesCommentService(repository)
// 		likesComments, err := service.GetAllLikesComments()

// 		assert.Equal(t, tt.expectedErrorLikesComment, err)
// 		assert.Equal(t, tt.expectedLikesComments, likesComments)
// 	}
// }
