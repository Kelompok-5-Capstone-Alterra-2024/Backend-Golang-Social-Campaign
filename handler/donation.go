package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DonationHandler struct {
	donationService service.DonationService
	userService     service.UserService
}

func NewDonationHandler(donationService service.DonationService, userService service.UserService) *DonationHandler {
	return &DonationHandler{donationService, userService}
}

func (h *DonationHandler) CreateDonation(c echo.Context) error {
	user_id, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	user, err := h.userService.GetUserByID(uint(user_id))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get user", err.Error()))
	}

	donationRequest := dto.DonationRequest{}
	c.Bind(&donationRequest)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid id", err.Error()))
	}
	donationRequest.ID = uint(id)
	donationRequest.User = user

	donation, err := h.donationService.CreateDonation(donationRequest)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create donation", err.Error()))
	}

	response := dto.ToDonationResponse(donation)

	return c.JSON(200, helper.ResponseWithData(true, "donation created successfully", response))

}

func (h *DonationHandler) GetUserDonation(c echo.Context) error {
	user_id, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 4
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	donations, err := h.donationService.GetDonationByUserID(c.Request().Context(), limit, offset, uint(user_id))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donations", err.Error()))
	}

	response := dto.ToAllDonationsResponses(donations)

	return c.JSON(200, helper.ResponseWithData(true, "donations retrieved successfully", response))
}

func (h *DonationHandler) GetDonationByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	donation, err := h.donationService.GetDonationByID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donation", err.Error()))
	}
	response := dto.ToHistoryDonationResponse(donation)
	return c.JSON(200, helper.ResponseWithData(true, "donation retrieved successfully", response))
}

func (h *DonationHandler) LikeComment(c echo.Context) error {
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	if err := h.donationService.LikeComment(c.Request().Context(), uint(commentID), uint(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *DonationHandler) UnLikeComment(c echo.Context) error {
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	if err := h.donationService.UnlikeComment(c.Request().Context(), uint(commentID), uint(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *DonationHandler) GetPaymentCallback(c echo.Context) error {
	var request dto.TransactionNotificationRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	err := h.donationService.PaymentProcess(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to process payment", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "payment processed successfully"))
}
