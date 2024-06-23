package service_test

// func TestCreateTransaction(t *testing.T) {

// 	testCases := []struct {
// 		title                    string
// 		mockTransaction          entities.Transaction
// 		mockErrorTransaction     error
// 		expectedTransaction      entities.Transaction
// 		expectedErrorTransaction error
// 	}{
// 		{
// 			title: "Success Create Transaction",
// 			mockTransaction: entities.Transaction{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        100000,
// 				BankName:      "BCA",
// 				NoRekening:    "1234567890",
// 				Name:          "Name",
// 				ImagePayment:  "image.jpg",
// 				FundraisingID: 1,
// 				Fundraising: entities.Fundraising{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Title:           "Title",
// 					Description:     "Description",
// 					Target:          1000000,
// 					CurrentProgress: 0,
// 					StartDate:       time.Now(),
// 					EndDate:         time.Now(),
// 					UserID:          1,
// 					User: entities.User{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Username: "Username",
// 						Email:    "user@gmail.com",
// 						Password: "password",
// 						Fullname: "Fullname",
// 						Role:     "user",
// 					},
// 				},
// 			},
// 			mockErrorTransaction: nil,
// 			expectedTransaction: entities.Transaction{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        100000,
// 				BankName:      "BCA",
// 				NoRekening:    "1234567890",
// 				Name:          "Name",
// 				ImagePayment:  "image.jpg",
// 				FundraisingID: 1,
// 				Fundraising: entities.Fundraising{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Title:           "Title",
// 					Description:     "Description",
// 					Target:          1000000,
// 					CurrentProgress: 0,
// 					StartDate:       time.Now(),
// 					EndDate:         time.Now(),
// 					UserID:          1,
// 					User: entities.User{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Username: "Username",
// 						Email:    "user@gmail.com",
// 						Password: "password",
// 						Fullname: "Fullname",
// 						Role:     "user",
// 					},
// 				},
// 			},
// 			expectedErrorTransaction: nil,
// 		},
// 		{
// 			title:                    "Failed Create Transaction",
// 			mockTransaction:          entities.Transaction{},
// 			mockErrorTransaction:     errors.New("Failed to create transaction"),
// 			expectedTransaction:      entities.Transaction{},
// 			expectedErrorTransaction: errors.New("Failed to create transaction"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewTransactionRepository(t)
// 		repository.On("Save", tt.mockTransaction).Return(tt.expectedTransaction, tt.mockErrorTransaction)

// 		service := service.NewTransactionService(repository)
// 		transaction, err := service.CreateTransaction(tt.mockTransaction)

// 		assert.Equal(t, tt.expectedErrorTransaction, err)
// 		assert.Equal(t, tt.expectedTransaction, transaction)
// 	}
// }

// func TestGetTransactionByID(t *testing.T) {

// 	testCases := []struct {
// 		title                    string
// 		mockTransaction          entities.Transaction
// 		mockErrorTransaction     error
// 		expectedTransaction      entities.Transaction
// 		expectedErrorTransaction error
// 	}{
// 		{
// 			title: "Success Get Transaction By ID",
// 			mockTransaction: entities.Transaction{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        100000,
// 				BankName:      "BCA",
// 				NoRekening:    "1234567890",
// 				Name:          "Name",
// 				ImagePayment:  "image.jpg",
// 				FundraisingID: 1,
// 				Fundraising: entities.Fundraising{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Title:           "Title",
// 					Description:     "Description",
// 					Target:          1000000,
// 					CurrentProgress: 0,
// 					StartDate:       time.Now(),
// 					EndDate:         time.Now(),
// 					UserID:          1,
// 					User: entities.User{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Username: "Username",
// 						Email:    "user@gmail.com",
// 					},
// 				},
// 			},
// 			mockErrorTransaction: nil,
// 			expectedTransaction: entities.Transaction{
// 				Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Amount:        100000,
// 				BankName:      "BCA",
// 				NoRekening:    "1234567890",
// 				Name:          "Name",
// 				ImagePayment:  "image.jpg",
// 				FundraisingID: 1,
// 				Fundraising: entities.Fundraising{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Title:           "Title",
// 					Description:     "Description",
// 					Target:          1000000,
// 					CurrentProgress: 0,
// 					StartDate:       time.Now(),
// 					EndDate:         time.Now(),
// 					UserID:          1,
// 					User: entities.User{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Username: "Username",
// 						Email:    "user@gmail.com",
// 					},
// 				},
// 			},
// 			expectedErrorTransaction: nil,
// 		},
// 		{
// 			title:                    "Failed Get Transaction By ID",
// 			mockTransaction:          entities.Transaction{},
// 			mockErrorTransaction:     errors.New("Transaction not found"),
// 			expectedTransaction:      entities.Transaction{},
// 			expectedErrorTransaction: errors.New("Transaction not found"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewTransactionRepository(t)
// 		repository.On("FindByID", tt.mockTransaction.ID).Return(tt.expectedTransaction, tt.mockErrorTransaction)

// 		service := service.NewTransactionService(repository)
// 		transaction, err := service.GetTransactionByID(tt.mockTransaction.ID)

// 		assert.Equal(t, tt.expectedErrorTransaction, err)
// 		assert.Equal(t, tt.expectedTransaction, transaction)
// 	}
// }

// func TestGetTransactions(t *testing.T) {

// 	testCases := []struct {
// 		title                     string
// 		mockTransactions          []entities.Transaction
// 		mockErrorTransactions     error
// 		expectedTransactions      []entities.Transaction
// 		expectedErrorTransactions error
// 	}{
// 		{
// 			title: "Success Get Transactions",
// 			mockTransactions: []entities.Transaction{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        100000,
// 					BankName:      "BCA",
// 					NoRekening:    "1234567890",
// 					Name:          "Name",
// 					ImagePayment:  "image.jpg",
// 					FundraisingID: 1,
// 					Fundraising: entities.Fundraising{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Title:           "Title",
// 						Description:     "Description",
// 						Target:          1000000,
// 						CurrentProgress: 0,
// 						StartDate:       time.Now(),
// 						EndDate:         time.Now(),
// 						UserID:          1,
// 						User: entities.User{
// 							Model: gorm.Model{
// 								ID: 1,
// 							},
// 							Username: "Username",
// 							Email:    "user@gmail.com",
// 						},
// 					},
// 				},
// 			},
// 			mockErrorTransactions: nil,
// 			expectedTransactions: []entities.Transaction{
// 				{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Amount:        100000,
// 					BankName:      "BCA",
// 					NoRekening:    "1234567890",
// 					Name:          "Name",
// 					ImagePayment:  "image.jpg",
// 					FundraisingID: 1,
// 					Fundraising: entities.Fundraising{
// 						Model: gorm.Model{
// 							ID: 1,
// 						},
// 						Title:           "Title",
// 						Description:     "Description",
// 						Target:          1000000,
// 						CurrentProgress: 0,
// 						StartDate:       time.Now(),
// 						EndDate:         time.Now(),
// 						UserID:          1,
// 						User: entities.User{
// 							Model: gorm.Model{
// 								ID: 1,
// 							},
// 							Username: "Username",
// 							Email:    "user@gmail.com",
// 						},
// 					},
// 				},
// 			},
// 			expectedErrorTransactions: nil,
// 		},
// 		{
// 			title:                     "Failed Get Transactions",
// 			mockTransactions:          []entities.Transaction{},
// 			mockErrorTransactions:     errors.New("Transactions not found"),
// 			expectedTransactions:      []entities.Transaction{},
// 			expectedErrorTransactions: errors.New("Transactions not found"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		repository := repositories.NewTransactionRepository(t)
// 		repository.On("FindAll", 1, 1).Return(tt.expectedTransactions, tt.mockErrorTransactions)

// 		service := service.NewTransactionService(repository)
// 		transactions, err := service.GetTransactions(1, 1)

// 		assert.Equal(t, tt.expectedErrorTransactions, err)
// 		assert.Equal(t, tt.expectedTransactions, transactions)
// 	}
// }
