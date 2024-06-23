package handler

import (
	"capstone/dto"
	"capstone/helper"
	middleware "capstone/middlewares"
	"capstone/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService        service.UserService
	fundraisingService service.FundraisingService
}

func NewUserHandler(userService service.UserService, fundraisingService service.FundraisingService) *UserHandler {
	return &UserHandler{userService, fundraisingService}
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

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(200, helper.ResponseWithData(true, "Token refreshed successfully", response))
}

func (h *UserHandler) ForgetPassword(c echo.Context) error {
	var request dto.ForgetPasswordRequest
	c.Bind(&request)
	err := h.userService.GenerateOTP(request.Email)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "OTP sent to your email"))
}

func (h *UserHandler) VerifyOTP(c echo.Context) error {
	type OTPRequest struct {
		OTP string `json:"otp"`
	}
	var request OTPRequest
	c.Bind(&request)
	_, err := h.userService.VerifyOTP(request.OTP)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", err.Error()))
	}
	return c.JSON(200, helper.ResponseWithData(true, "OTP verified successfully", request.OTP))
}

func (h *UserHandler) ResetPassword(c echo.Context) error {
	var request dto.ResetPasswordRequest
	c.Bind(&request)

	if request.Password != request.ConfirmPass {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", errors.New("password doesn't match").Error()))
	}

	err := h.userService.ResetPassword(request.OTP, request.Password)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid or expired OTP"))
	}
	return c.JSON(200, helper.ResponseWithData(true, "Password changed successfully", request.OTP))
}

func (h *UserHandler) ResetPasswordParamOtp(c echo.Context) error {
	otp := c.Param("otp")

	var request dto.ResetPasswordRequest
	c.Bind(&request)
	if request.Password != request.ConfirmPass {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", errors.New("password doesn't match").Error()))
	}

	err := h.userService.ResetPassword(otp, request.Password)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid or expired OTP"))
	}
	return c.JSON(200, helper.GeneralResponse(true, "Password changed successfully"))
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
		Fullname: userProfile.Fullname,
		Email:    userProfile.Email,
		NoTelp:   userProfile.NoTelp,
	}

	return c.JSON(200, helper.ResponseWithData(true, "Profile retrieved successfully", response))
}
func (h *UserHandler) EditProfile(c echo.Context) error {
	var request dto.EditProfileRequest
	c.Bind(&request)

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "Unauthorized.", err.Error()))
	}

	imgFile, err := c.FormFile("avatar_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid image url", err.Error()))
	}

	imageUrl, err := helper.UploadToCloudinary(imgFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to upload image", err.Error()))
	}

	request.Avatar = imageUrl
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
		NoTelp:   editProfile.NoTelp,
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

	return c.JSON(200, helper.ResponseWithData(true, "Success get history volunteer", history))
}

func (h *UserHandler) GetHistoryDonation(c echo.Context) error {
	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "Unauthorized.", err.Error()))
	}

	history, err := h.userService.GetHistoryDonation(uint(userID))
	if err != nil {
		return c.JSON(404, helper.ErrorResponse(false, "History not found.", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData(true, "", history))
}

func (h *UserHandler) CreateBookmarkFundraising(c echo.Context) error {
	fundraisingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	_, err = h.fundraisingService.FindFundraisingByID(fundraisingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to add bookmark", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	err = h.userService.AddUserFundraisingBookmark(uint(fundraisingID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to add bookmark", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "bookmark added successfully"))
}

func (h *UserHandler) DeleteBookmarkFundraising(c echo.Context) error {
	fundraisingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	err = h.userService.DeleteUserFundraisingBookmark(uint(fundraisingID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete bookmark", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "bookmark deleted successfully"))
}

func (h *UserHandler) GetBookmarkFundraising(c echo.Context) error {
	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 6
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	bookmarks, err := h.userService.GetUserFundraisingBookmark(uint(userID), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get bookmarks", err.Error()))
	}

	response := dto.ToAllUserFundraisingsResponse(bookmarks)

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "success get bookmarks", response, page, limit, int64(len(response))))
}

func (h *UserHandler) CreateBookmarkArticle(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	err = h.userService.AddUserArticleBookmark(uint(articleID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to add bookmark", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "bookmark added successfully"))
}

func (h *UserHandler) DeleteBookmarkArticle(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	err = h.userService.DeleteUserArticleBookmark(uint(articleID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete bookmark", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "bookmark deleted successfully"))
}

func (h *UserHandler) GetBookmarkArticle(c echo.Context) error {
	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 6
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	bookmarks, err := h.userService.GetUserArticleBookmark(uint(userID), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get bookmarks", err.Error()))
	}

	response := dto.ToAllUserArticleBookmarkResponse(bookmarks)

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "success get bookmarks", response, page, limit, int64(len(response))))
}

func (h *UserHandler) GetUserBookmarkVolunteer(c echo.Context) error {
	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 6
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	bookmarks, err := h.userService.GetUserVolunteerBookmark(uint(userID), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get bookmarks", err.Error()))
	}

	response := dto.ToAllUserVolunteerBookmarkResponse(bookmarks)

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "success get bookmarks", response, page, limit, int64(len(response))))
}

func (h *UserHandler) CreateBookmarkVolunteer(c echo.Context) error {
	volunteerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	err = h.userService.AddUserVolunteerBookmark(uint(volunteerID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to add bookmark", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "bookmark added successfully"))
}

func (h *UserHandler) DeleteBookmarkVolunteer(c echo.Context) error {
	volunteerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	err = h.userService.DeleteUserVolunteerBookmark(uint(volunteerID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete bookmark", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "bookmark deleted successfully"))
}

func (h *UserHandler) GetNotificationFundraising(c echo.Context) error {
	notifications, err := h.userService.GetNotificationFundraising()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get notifications", err.Error()))
	}

	response := dto.ToAllUserNotificationResponse(notifications)

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "success get notifications", response))
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
