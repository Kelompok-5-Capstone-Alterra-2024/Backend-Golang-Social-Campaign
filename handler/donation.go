package handler

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	"capstone/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DonationHandler struct {
	donationService service.DonationService
}

func NewDonationHandler(donationService service.DonationService) *DonationHandler {
	return &DonationHandler{donationService}
}

func (h *DonationHandler) CreateDonation(c echo.Context) error {
	user := c.Get("user").(entities.User)
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
	return c.JSON(200, helper.ResponseWithData(true, "donation created successfully", donation))

}
