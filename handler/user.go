package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Register(c echo.Context) error {
	var request dto.RegisterRequest
	c.Bind(&request)
	_, err := h.userService.Register(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse("failed", "validation failed", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse("success", "User registered successfully"))
}

func (h *UserHandler) Login(c echo.Context) error {
	var request dto.LoginRequest
	c.Bind(&request)
	loggedUser, err := h.userService.Login(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse("failed", "validation failed", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData("success", "User logged in successfully", loggedUser.Token))
}
