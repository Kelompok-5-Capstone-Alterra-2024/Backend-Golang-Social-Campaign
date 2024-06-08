package handler

import (
	"capstone/dto"
	"capstone/helper"
	middleware "capstone/middlewares"
	"capstone/service"
	"errors"
	"net/http"

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

// func (h *UserHandler) Login(c echo.Context) error {
// 	var request dto.LoginRequest
// 	c.Bind(&request)
// 	loggedUser, err := h.userService.Login(request)
// 	if err != nil {
// 		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid credentials"))
// 	}

// 	return c.JSON(200, helper.ResponseWithData(true, "User logged in successfully", loggedUser.Token))
// }

// func (h *UserHandler) Login(c echo.Context) error {
// 	var request dto.LoginRequest
// 	c.Bind(&request)
// 	_, accessToken, refreshToken, err := h.userService.Login(request)
// 	if err != nil {
// 		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid credentials"))
// 	}

// 	response := map[string]string{
// 		"access_token":  accessToken,
// 		"refresh_token": refreshToken,
// 	}
// 	return c.JSON(200, helper.ResponseWithData(true, "User logged in successfully", response))
// }

// handler/user.go
func (h *UserHandler) Login(c echo.Context) error {
	var request dto.LoginRequest
	c.Bind(&request)
	_, accessToken, refreshToken, err := h.userService.Login(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid credentials"))
	}

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(200, helper.ResponseWithData(true, "User logged in successfully", response))
}

func (h *UserHandler) RefreshToken(c echo.Context) error {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	c.Bind(&request)

	claims, err := middleware.VerifyRefreshToken(request.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid refresh token",
		})
	}

	accessToken, refreshToken := middleware.GenerateToken(claims.ID, claims.Username, claims.Role)

	return c.JSON(200, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
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

// func (h *UserHandler) RefreshToken(c echo.Context) error {
// 	refreshToken := c.Request().Header.Get("Refresh-Token")
// 	if refreshToken == "" {
// 		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 			"message": "Refresh token missing",
// 		})
// 	}

// 	newAccessToken, newRefreshToken, err := middleware.RefreshToken(refreshToken)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 			"message": "Invalid refresh token",
// 		})
// 	}

// 	response := map[string]string{
// 		"access_token":  newAccessToken,
// 		"refresh_token": newRefreshToken,
// 	}
// 	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "Token refreshed successfully", response))
// }
