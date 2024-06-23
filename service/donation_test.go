package service_test

// func TestCreateDonation(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockDonation          entities.Donation
// 		mockErrorDonation     error
// 		expectedDonation      entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title: "Success Create Donation",
// 			mockDonation: entities.Donation{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        10000,
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 				Code:          "12345",
// 				PaymentUrl:    "http://localhost:8080",
// 				PaymentMethod: "gopay",
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonation: entities.Donation{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        10000,
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 				Code:          "12345",
// 				PaymentUrl:    "http://localhost:8080",
// 				PaymentMethod: "gopay",
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Create Donation",
// 			mockDonation:          entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonation:      entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("Create", testCase.mockDonation).Return(testCase.expectedDonation, testCase.expectedErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donation, err := service.CreateDonation(testCase.mockDonation)

// 		assert.Equal(t, testCase.expectedDonation, donation)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestFetchStatusFromMidtrans(t *testing.T) {

// 	testCases := []struct {
// 		title          string
// 		mockOrderID    string
// 		mockError      error
// 		expectedResult entities.TransactionNotificationRequest
// 		expectedError  error
// 	}{
// 		{
// 			title:       "Success Fetch Status From Midtrans",
// 			mockOrderID: "12345",
// 			mockError:   nil,
// 			expectedResult: entities.TransactionNotificationRequest{
// 				TransactionStatus: "pending",
// 				OrderID:           "12345",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:          "Failed Fetch Status From Midtrans",
// 			mockOrderID:    "12345",
// 			mockError:      errors.New("Unexpected Error"),
// 			expectedResult: entities.TransactionNotificationRequest{},
// 			expectedError:  errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("FetchStatusFromMidtrans", testCase.mockOrderID).Return(testCase.expectedResult, testCase.expectedError)

// 		service := service.NewDonationService(repository, repository)

// 		result, err := service.FetchStatusFromMidtrans(testCase.mockOrderID)

// 		assert.Equal(t, testCase.expectedResult, result)
// 		assert.Equal(t, testCase.expectedError, err)
// 	}
// }

// func TestLikeDonationComment(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockLikeDonation      entities.LikeDonationComment
// 		mockErrorLikeDonation error
// 		expectedLikeDonation  entities.LikeDonationComment
// 		expectedError         error
// 	}{
// 		{
// 			title: "Success Like Donation Comment",
// 			mockLikeDonation: entities.LikeDonationComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationCommentID: 1,
// 				UserID:            1,
// 			},
// 			mockErrorLikeDonation: nil,
// 			expectedLikeDonation: entities.LikeDonationComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationCommentID: 1,
// 				UserID:            1,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:                 "Failed Like Donation Comment",
// 			mockLikeDonation:      entities.LikeDonationComment{},
// 			mockErrorLikeDonation: errors.New("Unexpected Error"),
// 			expectedLikeDonation:  entities.LikeDonationComment{},
// 			expectedError:         errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("LikeDonationComment", testCase.mockLikeDonation).Return(testCase.expectedLikeDonation, testCase.expectedError)

// 		service := service.NewDonationService(repository, repository)

// 		likeDonation, err := service.LikeDonationComment(testCase.mockLikeDonation)

// 		assert.Equal(t, testCase.expectedLikeDonation, likeDonation)
// 		assert.Equal(t, testCase.expectedError, err)
// 	}
// }

// func TestUnlikeDonationComment(t *testing.T) {

// 	testCases := []struct {
// 		title                   string
// 		mockLikeDonation        entities.LikeDonationComment
// 		mockErrorLikeDonation   error
// 		mockErrorUnlikeDonation error
// 		expectedError           error
// 	}{
// 		{
// 			title: "Success Unlike Donation Comment",
// 			mockLikeDonation: entities.LikeDonationComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationCommentID: 1,
// 				UserID:            1,
// 			},
// 			mockErrorLikeDonation:   nil,
// 			mockErrorUnlikeDonation: nil,
// 			expectedError:           nil,
// 		},
// 		{
// 			title: "Failed Unlike Donation Comment",
// 			mockLikeDonation: entities.LikeDonationComment{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				DonationCommentID: 1,
// 				UserID:            1,
// 			},
// 			mockErrorLikeDonation:   nil,
// 			mockErrorUnlikeDonation: errors.New("Unexpected Error"),
// 			expectedError:           errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("UnlikeDonationComment", testCase.mockLikeDonation).Return(testCase.mockErrorUnlikeDonation)

// 		service := service.NewDonationService(repository, repository)

// 		err := service.UnlikeDonationComment(testCase.mockLikeDonation)

// 		assert.Equal(t, testCase.expectedError, err)
// 	}
// }

// func TestGetDonationByUserID(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockLimit             int
// 		mockOffset            int
// 		mockUserID            uint
// 		mockDonations         []entities.Donation
// 		mockErrorDonations    error
// 		expectedDonations     []entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:      "Success Get Donation By User ID",
// 			mockLimit:  10,
// 			mockOffset: 0,
// 			mockUserID: 1,
// 			mockDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			mockErrorDonations: nil,
// 			expectedDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By User ID",
// 			mockLimit:             10,
// 			mockOffset:            0,
// 			mockUserID:            1,
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonations:    errors.New("Unexpected Error"),
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByUserID", testCase.mockLimit, testCase.mockOffset, testCase.mockUserID).Return(testCase.mockDonations, testCase.mockErrorDonations)

// 		service := service.NewDonationService(repository, repository)

// 		donations, err := service.GetDonationByUserID(testCase.mockLimit, testCase.mockOffset, testCase.mockUserID)

// 		assert.Equal(t, testCase.expectedDonations, donations)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationCommentByDonationID(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockDonationID        uint
// 		mockDonationComments  []entities.DonationComment
// 		mockErrorDonation     error
// 		expectedDonation      []entities.DonationComment
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:          "Success Get Donation Comment By Donation ID",
// 			mockDonationID: 1,
// 			mockDonationComments: []entities.DonationComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					DonationID: 1,
// 					Comment:    "Good Luck",
// 					TotalLikes: 0,
// 				},
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonation: []entities.DonationComment{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					DonationID: 1,
// 					Comment:    "Good Luck",
// 					TotalLikes: 0,
// 				},
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation Comment By Donation ID",
// 			mockDonationID:        1,
// 			mockDonationComments:  []entities.DonationComment{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonation:      []entities.DonationComment{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetCommentByDonationID", testCase.mockDonationID).Return(testCase.mockDonationComments, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donationComments, err := service.GetDonationCommentByDonationID(testCase.mockDonationID)

// 		assert.Equal(t, testCase.expectedDonation, donationComments)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationByID(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockDonationID        int
// 		mockDonation          entities.Donation
// 		mockErrorDonation     error
// 		expectedDonation      entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:          "Success Get Donation By ID",
// 			mockDonationID: 1,
// 			mockDonation: entities.Donation{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        10000,
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 				Code:          "12345",
// 				PaymentUrl:    "http://localhost:8080",
// 				PaymentMethod: "gopay",
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonation: entities.Donation{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        10000,
// 				UserID:        1,
// 				FundraisingID: 1,
// 				Status:        "pending",
// 				Code:          "12345",
// 				PaymentUrl:    "http://localhost:8080",
// 				PaymentMethod: "gopay",
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By ID",
// 			mockDonationID:        1,
// 			mockDonation:          entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonation:      entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByID", testCase.mockDonationID).Return(testCase.mockDonation, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donation, err := service.GetDonationByID(testCase.mockDonationID)

// 		assert.Equal(t, testCase.expectedDonation, donation)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationByFundraisingID(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockFundraisingID     int
// 		mockDonations         []entities.Donation
// 		mockErrorDonation     error
// 		expectedDonations     []entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:             "Success Get Donation By Fundraising ID",
// 			mockFundraisingID: 1,
// 			mockDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By Fundraising ID",
// 			mockFundraisingID:     1,
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByFundraisingID", testCase.mockFundraisingID).Return(testCase.mockDonations, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donations, err := service.GetDonationByFundraisingID(testCase.mockFundraisingID)

// 		assert.Equal(t, testCase.expectedDonations, donations)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationByFundraisingIDAndUserID(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockFundraisingID     int
// 		mockUserID            int
// 		mockDonations         []entities.Donation
// 		mockErrorDonation     error
// 		expectedDonations     []entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:             "Success Get Donation By Fundraising ID And User ID",
// 			mockFundraisingID: 1,
// 			mockUserID:        1,
// 			mockDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By Fundraising ID And User ID",
// 			mockFundraisingID:     1,
// 			mockUserID:            1,
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByFundraisingIDAndUserID", testCase.mockFundraisingID, testCase.mockUserID).Return(testCase.mockDonations, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donations, err := service.GetDonationByFundraisingIDAndUserID(testCase.mockFundraisingID, testCase.mockUserID)

// 		assert.Equal(t, testCase.expectedDonations, donations)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationByFundraisingIDAndStatus(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockFundraisingID     int
// 		mockStatus            string
// 		mockDonations         []entities.Donation
// 		mockErrorDonation     error
// 		expectedDonations     []entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:             "Success Get Donation By Fundraising ID And Status",
// 			mockFundraisingID: 1,
// 			mockStatus:        "pending",
// 			mockDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By Fundraising ID And Status",
// 			mockFundraisingID:     1,
// 			mockStatus:            "pending",
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByFundraisingIDAndStatus", testCase.mockFundraisingID, testCase.mockStatus).Return(testCase.mockDonations, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donations, err := service.GetDonationByFundraisingIDAndStatus(testCase.mockFundraisingID, testCase.mockStatus)

// 		assert.Equal(t, testCase.expectedDonations, donations)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationByFundraisingIDAndUserIDAndStatus(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockFundraisingID     int
// 		mockUserID            int
// 		mockStatus            string
// 		mockDonations         []entities.Donation
// 		mockErrorDonation     error
// 		expectedDonations     []entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:             "Success Get Donation By Fundraising ID And User ID And Status",
// 			mockFundraisingID: 1,
// 			mockUserID:        1,
// 			mockStatus:        "pending",
// 			mockDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			mockErrorDonation: nil,
// 			expectedDonations: []entities.Donation{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        10000,
// 					UserID:        1,
// 					FundraisingID: 1,
// 					Status:        "pending",
// 					Code:          "12345",
// 					PaymentUrl:    "http://localhost:8080",
// 					PaymentMethod: "gopay",
// 				},
// 			},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By Fundraising ID And User ID And Status",
// 			mockFundraisingID:     1,
// 			mockUserID:            1,
// 			mockStatus:            "pending",
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByFundraisingIDAndUserIDAndStatus", testCase.mockFundraisingID, testCase.mockUserID, testCase.mockStatus).Return(testCase.mockDonations, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donations, err := service.GetDonationByFundraisingIDAndUserIDAndStatus(testCase.mockFundraisingID, testCase.mockUserID, testCase.mockStatus)

// 		assert.Equal(t, testCase.expectedDonations, donations)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }

// func TestGetDonationByStatus(t *testing.T) {

// 	testCases := []struct {
// 		title                 string
// 		mockStatus            string
// 		mockDonations         []entities.Donation
// 		mockErrorDonation     error
// 		expectedDonations     []entities.Donation
// 		expectedErrorDonation error
// 	}{
// 		{
// 			title:                 "Success Get Donation By Status",
// 			mockStatus:            "pending",
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonation:     nil,
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: nil,
// 		},
// 		{
// 			title:                 "Failed Get Donation By Status",
// 			mockStatus:            "pending",
// 			mockDonations:         []entities.Donation{},
// 			mockErrorDonation:     errors.New("Unexpected Error"),
// 			expectedDonations:     []entities.Donation{},
// 			expectedErrorDonation: errors.New("Unexpected Error"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := repositories.NewDonationRepository(t)
// 		repository.On("GetByStatus", testCase.mockStatus).Return(testCase.mockDonations, testCase.mockErrorDonation)

// 		service := service.NewDonationService(repository, repository)

// 		donations, err := service.GetDonationByStatus(testCase.mockStatus)

// 		assert.Equal(t, testCase.expectedDonations, donations)
// 		assert.Equal(t, testCase.expectedErrorDonation, err)
// 	}
// }
