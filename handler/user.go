package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"errors"
	"strconv"

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

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "Unauthorized.", err.Error()))
	}

	userProfile, err := h.userService.GetUserProfile(userID)
	if err != nil {
		return c.JSON(404, helper.ErrorResponse(false, "Profile not found.", err.Error()))
	}

	response := dto.UserProfileResponse{
		ID:       userProfile.ID,
		Avatar:   userProfile.Avatar,
		Username: userProfile.Username,
	}

	return c.JSON(200, helper.ResponseWithData(true, "", response))
}

func (h *UserHandler) EditProfile(c echo.Context) error {
	var request dto.EditProfileRequest
	c.Bind(&request)

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "Unauthorized.", err.Error()))
	}

	// request.ID = uint(userID)
	editProfile, err := h.userService.EditProfile(userID, request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}

	response := dto.EditProfileRequest{
		// ID:       editProfile.ID,
		Fullname: editProfile.Fullname,
		Email:    editProfile.Email,
		Avatar:   editProfile.Avatar,
		Username: editProfile.Username,
	}

	return c.JSON(200, helper.ResponseWithData(true, "Profile updated successfully", response))
}

func (h *UserHandler) ChangePassword(c echo.Context) error {
	var request dto.ChangePasswordRequest
	c.Bind(&request)

	if request.NewPassword != request.ConfirmPassword {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", errors.New("password doesn't match").Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "Unauthorized.", err.Error()))
	}

	err = h.userService.ChangePassword(userID, request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "Password changed successfully"))
}

func (h *UserHandler) GetHistoryVolunteer(c echo.Context) error {
	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "Unauthorized.", err.Error()))
	}

	history, err := h.userService.GetHistoryVolunteer(uint(userID))
	if err != nil {
		return c.JSON(404, helper.ErrorResponse(false, "History not found.", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData(true, "", history))
}

func (h *UserHandler) GetHistoryVolunteerDetail(c echo.Context) error {
	historyID, _ := strconv.Atoi(c.Param("id"))
	history, err := h.userService.GetHistoryVolunteerDetail(historyID)
	if err != nil {
		return c.JSON(404, helper.ErrorResponse(false, "History not found.", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData(true, "", history))
}
