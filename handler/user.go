package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"errors"

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
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "User registered successfully"))
}

func (h *UserHandler) Login(c echo.Context) error {
	var request dto.LoginRequest
	c.Bind(&request)
	loggedUser, err := h.userService.Login(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid credentials"))
	}

	return c.JSON(200, helper.ResponseWithData(true, "User logged in successfully", loggedUser.Token))
}

func (h *UserHandler) ForgetPassword(c echo.Context) error {
	var request dto.ForgetPasswordRequest
	c.Bind(&request)
	err := h.userService.GenerateResetToken(request.Email)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "Reset password link sent to your email"))
}

func (h *UserHandler) ResetPassword(c echo.Context) error {
	var request dto.ResetPasswordRequest
	c.Bind(&request)

	if request.Password != request.ConfirmPass {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", errors.New("password doesn't match").Error()))
	}

	resetToken := c.QueryParam("token")
	if resetToken == "" {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "token not found"))
	}

	err := h.userService.ResetPassword(resetToken, request.Password)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "Password reset successfully"))
}
