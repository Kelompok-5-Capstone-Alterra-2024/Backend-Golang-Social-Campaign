package service_test

// import (
// 	"capstone/dto"
// 	"capstone/entities"
// 	"capstone/mocks/repositories"
// 	"capstone/service"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestRegister(t *testing.T) {
// 	testCases := []struct {
// 		title            string
// 		request          dto.RegisterRequest
// 		mockFindEmail    entities.User
// 		mockFindUsername entities.User
// 		mockFindNoTelp   entities.User
// 		mockFindFullName entities.User
// 		mockSaveError    error
// 		expectedError    error
// 	}{
// 		{
// 			title: "Register Success",
// 			request: dto.RegisterRequest{
// 				Fullname:    "Test User",
// 				Username:    "testuser",
// 				Email:       "testuser@example.com",
// 				Password:    "password",
// 				ConfirmPass: "password",
// 				NoTelp:      "1234567890",
// 			},
// 			mockFindEmail:    entities.User{},
// 			mockFindUsername: entities.User{},
// 			mockFindNoTelp:   entities.User{},
// 			mockFindFullName: entities.User{},
// 			mockSaveError:    nil,
// 			expectedError:    nil,
// 		},
// 		{
// 			title: "Register Password Mismatch",
// 			request: dto.RegisterRequest{
// 				Fullname:    "Test User",
// 				Username:    "testuser",
// 				Email:       "testuser@example.com",
// 				Password:    "password",
// 				ConfirmPass: "password1",
// 				NoTelp:      "1234567890",
// 			},
// 			mockFindEmail:    entities.User{},
// 			mockFindUsername: entities.User{},
// 			mockFindNoTelp:   entities.User{},
// 			mockFindFullName: entities.User{},
// 			mockSaveError:    nil,
// 			expectedError:    errors.New("password doesn't match"),
// 		},
// 		{
// 			title: "Register Email Exists",
// 			request: dto.RegisterRequest{
// 				Fullname:    "Test User",
// 				Username:    "testuser",
// 				Email:       "testuser@example.com",
// 				Password:    "password",
// 				ConfirmPass: "password",
// 				NoTelp:      "1234567890",
// 			},
// 			mockFindEmail: entities.User{
// 				Email: "testuser@example.com",
// 			},
// 			mockFindUsername: entities.User{},
// 			mockFindNoTelp:   entities.User{},
// 			mockFindFullName: entities.User{},
// 			mockSaveError:    nil,
// 			expectedError:    errors.New("email already exists"),
// 		},
// 		{
// 			title: "Register Username Exists",
// 			request: dto.RegisterRequest{
// 				Fullname:    "Test User",
// 				Username:    "testuser",
// 				Email:       "testuser@example.com",
// 				Password:    "password",
// 				ConfirmPass: "password",
// 				NoTelp:      "1234567890",
// 			},
// 			mockFindEmail: entities.User{},
// 			mockFindUsername: entities.User{
// 				Username: "testuser",
// 			},
// 			mockFindNoTelp:   entities.User{},
// 			mockFindFullName: entities.User{},
// 			mockSaveError:    nil,
// 			expectedError:    errors.New("username already exists"),
// 		},
// 		{
// 			title: "Register Phone Number Exists",
// 			request: dto.RegisterRequest{
// 				Fullname:    "Test User",
// 				Username:    "testuser",
// 				Email:       "testuser@example.com",
// 				Password:    "password",
// 				ConfirmPass: "password",
// 				NoTelp:      "1234567890",
// 			},
// 			mockFindEmail:    entities.User{},
// 			mockFindUsername: entities.User{},
// 			mockFindNoTelp: entities.User{
// 				NoTelp: "1234567890",
// 			},
// 			mockFindFullName: entities.User{},
// 			mockSaveError:    nil,
// 			expectedError:    errors.New("phone number already exists"),
// 		},
// 		{
// 			title: "Register Fullname Exists",
// 			request: dto.RegisterRequest{
// 				Fullname:    "Test User",
// 				Username:    "testuser",
// 				Email:       "testuser@example.com",
// 				Password:    "password",
// 				ConfirmPass: "password",
// 				NoTelp:      "1234567890",
// 			},
// 			mockFindEmail:    entities.User{},
// 			mockFindUsername: entities.User{},
// 			mockFindNoTelp:   entities.User{},
// 			mockFindFullName: entities.User{
// 				Fullname: "Test User",
// 			},
// 			mockSaveError: nil,
// 			expectedError: errors.New("fullname already exists"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("FindByEmail", testCase.request.Email).Return(testCase.mockFindEmail, nil)
// 		repository.On("FindByUsername", testCase.request.Username).Return(testCase.mockFindUsername, nil)
// 		repository.On("FindByNoTelp", testCase.request.NoTelp).Return(testCase.mockFindNoTelp, nil)
// 		repository.On("FindByFullName", testCase.request.Fullname).Return(testCase.mockFindFullName, nil)

// 		if testCase.expectedError == nil {
// 			repository.On("Save", mock.Anything).Return(testCase.request, testCase.mockSaveError)
// 		}

// 		service := service.NewUserService(repository)
// 		user, err := service.Register(testCase.request)

// 		assert.Equal(t, testCase.expectedError, err)
// 		if testCase.expectedError == nil {
// 			assert.Equal(t, testCase.request.Fullname, user.Fullname)
// 			assert.Equal(t, testCase.request.Username, user.Username)
// 			assert.Equal(t, testCase.request.Email, user.Email)
// 			assert.Equal(t, testCase.request.NoTelp, user.NoTelp)
// 		}
// 	}
// }

// func TestLogin(t *testing.T) {
// 	testCases := []struct {
// 		title         string
// 		request       dto.LoginRequest
// 		mockUser      entities.User
// 		mockError     error
// 		expectedError error
// 	}{
// 		{
// 			title: "Login Success",
// 			request: dto.LoginRequest{
// 				Username: "testuser",
// 				Password: "password",
// 			},
// 			mockUser: entities.User{
// 				Username: "testuser",
// 				Password: "password",
// 			},
// 			mockError:     nil,
// 			expectedError: nil,
// 		},
// 		{
// 			title: "Login Wrong Password",
// 			request: dto.LoginRequest{
// 				Username: "testuser",
// 				Password: "wrongpassword",
// 			},
// 			mockUser: entities.User{
// 				Username: "testuser",
// 				Password: "password",
// 			},
// 			mockError:     nil,
// 			expectedError: errors.New("wrong password"),
// 		},
// 		{
// 			title: "Login User Not Found",
// 			request: dto.LoginRequest{
// 				Username: "nonexistentuser",
// 				Password: "password",
// 			},
// 			mockUser:      entities.User{},
// 			mockError:     errors.New("user not found"),
// 			expectedError: errors.New("user not found"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("FindByUsername", testCase.request.Username).Return(testCase.mockUser, testCase.mockError)

// 		service := service.NewUserService(repository)
// 		user, accessToken, refreshToken, err := service.Login(testCase.request)

// 		assert.Equal(t, testCase.expectedError, err)
// 		if testCase.expectedError == nil {
// 			assert.Equal(t, testCase.mockUser.Username, user.Username)
// 			assert.NotEmpty(t, accessToken)
// 			assert.NotEmpty(t, refreshToken)
// 		}
// 	}
// }

// func TestGenerateResetToken(t *testing.T) {
// 	testCases := []struct {
// 		title         string
// 		email         string
// 		mockUser      entities.User
// 		mockError     error
// 		expectedError error
// 	}{
// 		{
// 			title: "Generate Reset Token Success",
// 			email: "testuser@example.com",
// 			mockUser: entities.User{
// 				Email: "testuser@example.com",
// 			},
// 			mockError:     nil,
// 			expectedError: nil,
// 		},
// 		{
// 			title:         "Generate Reset Token User Not Found",
// 			email:         "nonexistent@example.com",
// 			mockUser:      entities.User{},
// 			mockError:     errors.New("user not found"),
// 			expectedError: errors.New("user not found"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("FindByEmail", testCase.email).Return(testCase.mockUser, testCase.mockError)

// 		service := service.NewUserService(repository)
// 		err := service.GenerateResetToken(testCase.email)

// 		assert.Equal(t, testCase.expectedError, err)
// 	}
// }

// func TestResetPassword(t *testing.T) {
// 	testCases := []struct {
// 		title         string
// 		otp           string
// 		newPassword   string
// 		mockUser      entities.User
// 		mockError     error
// 		expectedError error
// 	}{
// 		{
// 			title:       "Reset Password Success",
// 			otp:         "123456",
// 			newPassword: "newpassword",
// 			mockUser: entities.User{
// 				Otp: "123456",
// 			},
// 			mockError:     nil,
// 			expectedError: nil,
// 		},
// 		{
// 			title:         "Reset Password OTP Mismatch",
// 			otp:           "wrongotp",
// 			newPassword:   "newpassword",
// 			mockUser:      entities.User{},
// 			mockError:     errors.New("invalid OTP"),
// 			expectedError: errors.New("invalid OTP"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("FindByOtp", testCase.otp).Return(testCase.mockUser, testCase.mockError)
// 		if testCase.expectedError == nil {
// 			repository.On("UpdatePassword", testCase.mockUser, testCase.newPassword).Return(nil)
// 		}

// 		service := service.NewUserService(repository)
// 		err := service.ResetPassword(testCase.otp, testCase.newPassword)

// 		assert.Equal(t, testCase.expectedError, err)
// 	}
// }

// func TestDeleteUser(t *testing.T) {
// 	testCases := []struct {
// 		title         string
// 		userId        uint
// 		mockError     error
// 		expectedError error
// 	}{
// 		{
// 			title:         "Delete User Success",
// 			userId:        1,
// 			mockError:     nil,
// 			expectedError: nil,
// 		},
// 		{
// 			title:         "Delete User Not Found",
// 			userId:        999,
// 			mockError:     errors.New("user not found"),
// 			expectedError: errors.New("user not found"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("DeleteUser", testCase.userId).Return(testCase.mockError)

// 		service := service.NewUserService(repository)
// 		err := service.DeleteUser(testCase.userId)

// 		assert.Equal(t, testCase.expectedError, err)
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	testCases := []struct {
// 		title         string
// 		userId        uint
// 		mockUser      entities.User
// 		mockError     error
// 		expectedUser  entities.User
// 		expectedError error
// 	}{
// 		{
// 			title:  "Get User Success",
// 			userId: 1,
// 			mockUser: entities.User{
// 				ID:       1,
// 				Username: "testuser",
// 			},
// 			mockError: nil,
// 			expectedUser: entities.User{
// 				ID:       1,
// 				Username: "testuser",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:         "Get User Not Found",
// 			userId:        999,
// 			mockUser:      entities.User{},
// 			mockError:     errors.New("user not found"),
// 			expectedUser:  entities.User{},
// 			expectedError: errors.New("user not found"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("FindByID", testCase.userId).Return(testCase.mockUser, testCase.mockError)

// 		service := service.NewUserService(repository)
// 		user, err := service.GetUser(testCase.userId)

// 		assert.Equal(t, testCase.expectedError, err)
// 		assert.Equal(t, testCase.expectedUser, user)
// 	}
// }

// func TestUpdateUser(t *testing.T) {
// 	testCases := []struct {
// 		title          string
// 		userId         uint
// 		updateRequest  dto.UpdateUserRequest
// 		mockFindUser   entities.User
// 		mockUpdateUser entities.User
// 		mockError      error
// 		expectedUser   entities.User
// 		expectedError  error
// 	}{
// 		{
// 			title:  "Update User Success",
// 			userId: 1,
// 			updateRequest: dto.UpdateUserRequest{
// 				Fullname: "Updated User",
// 				NoTelp:   "0987654321",
// 			},
// 			mockFindUser: entities.User{
// 				ID:       1,
// 				Fullname: "Old User",
// 				NoTelp:   "1234567890",
// 			},
// 			mockUpdateUser: entities.User{
// 				ID:       1,
// 				Fullname: "Updated User",
// 				NoTelp:   "0987654321",
// 			},
// 			mockError: nil,
// 			expectedUser: entities.User{
// 				ID:       1,
// 				Fullname: "Updated User",
// 				NoTelp:   "0987654321",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			title:  "Update User Not Found",
// 			userId: 999,
// 			updateRequest: dto.UpdateUserRequest{
// 				Fullname: "Updated User",
// 				NoTelp:   "0987654321",
// 			},
// 			mockFindUser:   entities.User{},
// 			mockUpdateUser: entities.User{},
// 			mockError:      errors.New("user not found"),
// 			expectedUser:   entities.User{},
// 			expectedError:  errors.New("user not found"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		repository := new(repositories.UserRepository)
// 		repository.On("FindByID", testCase.userId).Return(testCase.mockFindUser, testCase.mockError)
// 		if testCase.expectedError == nil {
// 			repository.On("UpdateUser", mock.AnythingOfType("entities.User")).Return(testCase.mockUpdateUser, testCase.mockError)
// 		}

// 		service := service.NewUserService(repository)
// 		user, err := service.UpdateUser(testCase.userId, testCase.updateRequest)

// 		assert.Equal(t, testCase.expectedError, err)
// 		if testCase.expectedError == nil {
// 			assert.Equal(t, testCase.expectedUser, user)
// 		}
// 	}
// }
