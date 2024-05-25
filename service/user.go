package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/repositories"
	"fmt"
)

type UserService interface {
	Register(request dto.RegisterRequest) (entities.User, error)
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

	user := entities.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Avatar:   "https://res.cloudinary.com/dvrhf8d9t/image/upload/v1715517059/default-avatar_yt6eua.png",
	}

	return s.userRepository.Save(user)
}
