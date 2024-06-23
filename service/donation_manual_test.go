package service_test

// func TestCreateManualDonation(t *testing.T) {

// 	testCases := []struct {
// 		title            string
// 		mockDonation     entities.DonationManual
// 		mockError        error
// 		expectedDonation entities.DonationManual
// 		expectedError    error
// 	}{
// 		{
// 			title: "Success Create Manual Donation",
// 			mockDonation: entities.DonationManual{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				ImagePayment:  "image.jpg",
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 			},
// 			mockError: nil,
// 			expectedDonation: entities.DonationManual{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				ImagePayment:  "image.jpg",
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:            "Failed Create Manual Donation",
// 			mockDonation:     entities.DonationManual{},
// 			mockError:        errors.New("Failed to create manual donation"),
// 			expectedDonation: entities.DonationManual{},
// 			expectedError:    errors.New("Failed to create manual donation"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("Save", tt.mockDonation).Return(tt.expectedDonation, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		donation, err := service.CreateManualDonation(tt.mockDonation)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedDonation, donation)
// 	}
// }

// func TestGetDonationManualByID(t *testing.T) {

// 	testCases := []struct {
// 		title            string
// 		mockDonation     entities.DonationManual
// 		mockError        error
// 		expectedDonation entities.DonationManual
// 		expectedError    error
// 	}{
// 		{
// 			title: "Success Get Donation Manual By ID",
// 			mockDonation: entities.DonationManual{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				ImagePayment:  "image.jpg",
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 			},
// 			mockError: nil,
// 			expectedDonation: entities.DonationManual{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				ImagePayment:  "image.jpg",
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:            "Failed Get Donation Manual By ID",
// 			mockDonation:     entities.DonationManual{},
// 			mockError:        errors.New("Failed to get donation manual"),
// 			expectedDonation: entities.DonationManual{},
// 			expectedError:    errors.New("Failed to get donation manual"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("GetByID", 1).Return(tt.mockDonation, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		donation, err := service.GetDonationManualByID(1)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedDonation, donation)
// 	}
// }

// func TestGetDonationManualByUserID(t *testing.T) {

// 	testCases := []struct {
// 		title             string
// 		mockDonations     []entities.DonationManual
// 		mockError         error
// 		expectedDonations []entities.DonationManual
// 		expectedError     error
// 	}{
// 		{
// 			title: "Success Get Donation Manual By User ID",
// 			mockDonations: []entities.DonationManual{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					ImagePayment:  "image.jpg",
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 				},
// 			},
// 			mockError: nil,
// 			expectedDonations: []entities.DonationManual{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					ImagePayment:  "image.jpg",
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 				},
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:             "Failed Get Donation Manual By User ID",
// 			mockDonations:     []entities.DonationManual{},
// 			mockError:         errors.New("Failed to get donation manual"),
// 			expectedDonations: []entities.DonationManual{},
// 			expectedError:     errors.New("Failed to get donation manual"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("GetByUserID", 1, 0, 1).Return(tt.mockDonations, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		donations, err := service.GetDonationManualByUserID(1, 0, 1)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedDonations, donations)
// 	}
// }

// func TestGetDonationCommentByFundraisingID(t *testing.T) {

// 	testCases := []struct {
// 		title            string
// 		mockDonation     []entities.DonationManualComment
// 		mockError        error
// 		expectedDonation []entities.DonationManualComment
// 		expectedError    error
// 	}{
// 		{
// 			title: "Success Get Donation Comment By Fundraising ID",
// 			mockDonation: []entities.DonationManualComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Comment:          "Comment",
// 					DonationManualID: 1,
// 					TotalLikes:       0,
// 				},
// 			},
// 			mockError: nil,
// 			expectedDonation: []entities.DonationManualComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Comment:          "Comment",
// 					DonationManualID: 1,
// 					TotalLikes:       0,
// 				},
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:            "Failed Get Donation Comment By Fundraising ID",
// 			mockDonation:     []entities.DonationManualComment{},
// 			mockError:        errors.New("Failed to get donation comment"),
// 			expectedDonation: []entities.DonationManualComment{},
// 			expectedError:    errors.New("Failed to get donation comment"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("GetCommentByFundraisingID", 1).Return(tt.mockDonation, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		donation, err := service.GetDonationCommentByFundraisingID(1)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedDonation, donation)
// 	}
// }

// func TestLikeComment(t *testing.T) {

// 	testCases := []struct {
// 		title         string
// 		mockLike      entities.LikeDonationManualComment
// 		mockError     error
// 		expectedLike  entities.LikeDonationManualComment
// 		expectedError error
// 	}{
// 		{
// 			title: "Success Like Comment",
// 			mockLike: entities.LikeDonationManualComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationManualCommentID: 1,
// 				UserID:                  1,
// 			},
// 			mockError: nil,
// 			expectedLike: entities.LikeDonationManualComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationManualCommentID: 1,
// 				UserID:                  1,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:         "Failed Like Comment",
// 			mockLike:      entities.LikeDonationManualComment{},
// 			mockError:     errors.New("Failed to like comment"),
// 			expectedLike:  entities.LikeDonationManualComment{},
// 			expectedError: errors.New("Failed to like comment"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("LikeComment", tt.mockLike).Return(tt.expectedLike, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		like, err := service.LikeComment(tt.mockLike)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedLike, like)
// 	}
// }

// func UnlikeComment(t *testing.T) {

// 	testCases := []struct {
// 		title         string
// 		mockLike      entities.LikeDonationManualComment
// 		mockError     error
// 		expectedLike  entities.LikeDonationManualComment
// 		expectedError error
// 	}{
// 		{
// 			title: "Success Unlike Comment",
// 			mockLike: entities.LikeDonationManualComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationManualCommentID: 1,
// 				UserID:                  1,
// 			},
// 			mockError: nil,
// 			expectedLike: entities.LikeDonationManualComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationManualCommentID: 1,
// 				UserID:                  1,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:         "Failed Unlike Comment",
// 			mockLike:      entities.LikeDonationManualComment{},
// 			mockError:     errors.New("Failed to unlike comment"),
// 			expectedLike:  entities.LikeDonationManualComment{},
// 			expectedError: errors.New("Failed to unlike comment"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("UnlikeComment", tt.mockLike).Return(tt.expectedLike, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		like, err := service.UnlikeComment(tt.mockLike)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedLike, like)
// 	}
// }

// func TestGetByFundraisingID(t *testing.T) {

// 	testCases := []struct {
// 		title            string
// 		mockDonation     []entities.DonationManual
// 		mockError        error
// 		expectedDonation []entities.DonationManual
// 		expectedError    error
// 	}{
// 		{
// 			title: "Success Get By Fundraising ID",
// 			mockDonation: []entities.DonationManual{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					ImagePayment:  "image.jpg",
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 				},
// 			},
// 			mockError: nil,
// 			expectedDonation: []entities.DonationManual{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					ImagePayment:  "image.jpg",
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 				},
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:            "Failed Get By Fundraising ID",
// 			mockDonation:     []entities.DonationManual{},
// 			mockError:        errors.New("Failed to get donation manual"),
// 			expectedDonation: []entities.DonationManual{},
// 			expectedError:    errors.New("Failed to get donation manual"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewDonationManualRepository(t)
// 		repository.On("GetByFundraisingID", 1).Return(tt.mockDonation, tt.mockError)

// 		service := service.NewDonationManualService(repository, nil)
// 		donation, err := service.GetByFundraisingID(1)

// 		assert.Equal(t, tt.expectedError, err)
// 		assert.Equal(t, tt.expectedDonation, donation)
// 	}
// }
