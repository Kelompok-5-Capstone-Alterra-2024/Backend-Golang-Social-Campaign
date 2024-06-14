package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"errors"
	"fmt"
)

type UserService interface {
	Register(request dto.RegisterRequest) (entities.User, error)
	Login(request dto.LoginRequest) (entities.User, error)
	GetUserByID(id uint) (entities.User, error)
	GenerateResetToken(email string) error
	ResetPassword(resetToken, newPassword string) error
	GenerateOTP(email string) error
	GetUserProfile(id int) (entities.User, error)
	EditProfile(userid int, request dto.EditProfileRequest) (entities.User, error)
	ChangePassword(userid int, request dto.ChangePasswordRequest) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) Register(request dto.RegisterRequest) (entities.User, error) {

	if request.ConfirmPass != request.Password {
		return entities.User{}, fmt.Errorf("password doesn't match")
	}

	userEmail, _ := s.userRepository.FindByEmail(request.Email)
	if userEmail.Email == request.Email {
		return entities.User{}, errors.New("email already exists")
	}

	userName, _ := s.userRepository.FindByUsername(request.Username)
	if userName.Username == request.Username {
		return entities.User{}, errors.New("username already exists")
	}

	userTelp, _ := s.userRepository.FindByNoTelp(request.NoTelp)
	if userTelp.NoTelp == request.NoTelp {
		return entities.User{}, errors.New("phone number already exists")
	}

	userFull, _ := s.userRepository.FindByFullName(request.Fullname)
	if userFull.Fullname == request.Fullname {
		return entities.User{}, errors.New("fullname already exists")
	}

	user := entities.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		NoTelp:   request.NoTelp,
		Avatar:   "https://res.cloudinary.com/dvrhf8d9t/image/upload/v1715517059/default-avatar_yt6eua.png",
	}

	return s.userRepository.Save(user)
}

func (s *userService) Login(request dto.LoginRequest) (entities.User, error) {
	username := request.Username
	password := request.Password

	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.Password != password {
		return user, fmt.Errorf("wrong password")
	}

	user.Token = middleware.GenerateToken(user.ID, user.Username, "user")

	return user, nil
}

func (s *userService) GenerateResetToken(email string) error {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	resetToken := helper.GenerateToken()
	user.ResetToken = resetToken
	err = s.userRepository.Update(user)
	if err != nil {
		return err
	}

	return helper.SendTokenRestPassword(email, resetToken)
}

func (s *userService) GenerateOTP(email string) error {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	otp := helper.GenerateRandomOTP(6)
	user.OTP = otp

	err = s.userRepository.Update(user)
	if err != nil {
		return err
	}

	return helper.SendOtpResetPassword(email, otp)
}

func (s *userService) ResetPassword(otp, newPassword string) error {
	// user, err := s.userRepository.FindByResetToken(resetToken)
	// if err != nil {
	// 	return err
	// }

	// if user.ResetToken != resetToken {
	// 	return errors.New("invalid or expired reset token")
	// }

	// user.Password = newPassword
	// user.ResetToken = ""
	user, err := s.userRepository.FindByOTP(otp)
	if err != nil {
		return err
	}

	if user.OTP != otp {
		return errors.New("invalid or expired otp")
	}

	user.OTP = ""
	user.Password = newPassword

	return s.userRepository.Update(user)
}

func (s *userService) GetUserByID(id uint) (entities.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) GetUserProfile(id int) (entities.User, error) {
	return s.userRepository.FindByID(uint(id))
}

func (s *userService) EditProfile(userid int, request dto.EditProfileRequest) (entities.User, error) {

	var existingUser entities.User
	existingUser, err := s.userRepository.FindByID(uint(userid))
	if err != nil {
		return entities.User{}, err
	}

	if request.Fullname != "" {
		existingUser.Fullname = request.Fullname
	}
	if request.Username != "" {
		existingUser.Username = request.Username
	}
	if request.Avatar != "" {
		existingUser.Avatar = request.Avatar
	}
	if request.Email != "" {
		existingUser.Email = request.Email
	}

	return existingUser, s.userRepository.UpdateProfile(uint(userid), existingUser)
}

func (s *userService) ChangePassword(userid int, request dto.ChangePasswordRequest) error {
	user, err := s.userRepository.FindByID(uint(userid))
	if err != nil {
		return err
	}

	if user.Password != request.CurrentPassword {
		return errors.New("wrong current password")
	}

	if request.NewPassword != request.ConfirmPassword {
		return errors.New("password doesn't match")
	}

	user.Password = request.NewPassword
	return s.userRepository.Update(user)
}
