package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"errors"
	"fmt"
	"time"
)

type UserService interface {
	Register(request dto.RegisterRequest) (entities.User, error)
	Login(request dto.LoginRequest) (entities.User, error)
	GetUserByID(id uint) (entities.User, error)
	GenerateResetToken(email string) error
	ResetPassword(resetToken, newPassword string) error
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

	userDB, _ := s.userRepository.FindByEmail(request.Email)
	if userDB.Email == request.Email {
		return entities.User{}, errors.New("email already exists")
	}

	user := entities.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
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
	user.ResetTokenExpire = time.Now().Add(1 * time.Hour)
	err = s.userRepository.Update(user)
	if err != nil {
		return err
	}

	return helper.SendTokenRestPassword(email, resetToken)
}

func (s *userService) ResetPassword(resetToken, newPassword string) error {
	user, err := s.userRepository.FindByResetToken(resetToken)
	if err != nil {
		return err
	}

	if user.ResetToken != resetToken {
		return errors.New("invalid or expired reset token")
	}

	user.Password = newPassword
	user.ResetToken = ""
	user.ResetTokenExpire = time.Time{}
	return s.userRepository.Update(user)
}

func (s *userService) GetUserByID(id uint) (entities.User, error) {
	return s.userRepository.FindByID(id)
}
