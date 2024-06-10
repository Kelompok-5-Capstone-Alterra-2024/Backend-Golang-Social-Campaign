package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DonationManualHandler struct {
	donationManualService service.DonationManualService
	service.UserService
}

func NewDonationManualHandler(donationManualService service.DonationManualService, userService service.UserService) *DonationManualHandler {
	return &DonationManualHandler{donationManualService, userService}
}

func (h *DonationManualHandler) CreateManualDonation(c echo.Context) error {
	userId, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	user, err := h.UserService.GetUserByID(uint(userId))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get user", err.Error()))
	}

	donationManual := dto.ManualDonationRequest{}
	c.Bind(&donationManual)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid fundraising id", err.Error()))
	}

	donationManual.ID = uint(id)
	donationManual.User = user

	donation, err := h.donationManualService.CreateManualDonation(donationManual)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create donation", err.Error()))
	}

	response := dto.ToManualDonationResponse(donation)

	return c.JSON(200, helper.ResponseWithData(true, "donation created successfully", response))
}

func (h *DonationManualHandler) GetDonationManualByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid id", err.Error()))
	}
	donation, err := h.donationManualService.GetDonationManualByID(id)

	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donation", err.Error()))
	}

	response := dto.ToHistoryDonationManualResponse(donation)
	return c.JSON(200, helper.ResponseWithData(true, "donation retrieved successfully", response))
}

func (h *DonationManualHandler) GetDonationManualByUserID(c echo.Context) error {
	userId, err := helper.GetUserIDFromJWT(c)
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
	donations, err := h.donationManualService.GetDonationManualByUserID(limit, offset, uint(userId))

	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donations", err.Error()))
	}

	response := dto.ToDonationManualsResponses(donations)
	return c.JSON(200, helper.ResponseWithData(true, "donations retrieved successfully", response))
}

func (h *DonationManualHandler) LikeComment(c echo.Context) error {
	userId, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid comment id", err.Error()))
	}
	err = h.donationManualService.LikeComment(c.Request().Context(), uint(commentId), uint(userId))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to like comment", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "comment liked successfully"))
}

func (h *DonationManualHandler) UnlikeComment(c echo.Context) error {
	userId, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid comment id", err.Error()))
	}
	err = h.donationManualService.UnlikeComment(c.Request().Context(), uint(commentId), uint(userId))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to unlike comment", err.Error()))
	}
	return c.JSON(200, helper.GeneralResponse(true, "comment unliked successfully"))
}
