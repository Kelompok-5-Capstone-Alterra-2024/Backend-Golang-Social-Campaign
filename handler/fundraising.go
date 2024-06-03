package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FundraisingHandler struct {
	fundraisingService service.FundraisingService
	donationService    service.DonationService
}

func NewFundraisingHandler(fundraisingService service.FundraisingService, donationService service.DonationService) *FundraisingHandler {
	return &FundraisingHandler{fundraisingService, donationService}
}

func (h *FundraisingHandler) GetFundraisings(c echo.Context) error {
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

	fundraisings, err := h.fundraisingService.FindFundraisings(c.Request().Context(), limit, offset)

	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraisings", err.Error()))
	}

	response := dto.ToAllFundraisingsResponse(fundraisings)

	return c.JSON(200, helper.ResponseWithData(true, "fundraisings retrieved successfully", response))

}

func (h *FundraisingHandler) GetFundraisingByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fundraising, err := h.fundraisingService.FindFundraisingByID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraising", err.Error()))
	}

	comments, err := h.donationService.GetDonationCommentByFundraisingID(id)

	donations, err := h.donationService.GetByFundraisingID(id)

	response := dto.ToFundraisingResponse(fundraising, comments, donations)
	return c.JSON(200, helper.ResponseWithData(true, "fundraising retrieved successfully", response))
}

func (h *FundraisingHandler) GetAllFundraisingCategories(c echo.Context) error {
	fundraisingCategories, err := h.fundraisingService.FindAllFundraisingCategories()
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraising categories", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData(true, "fundraising categories retrieved successfully", fundraisingCategories))
}

func (h *FundraisingHandler) GetFundraisingsByCategoryID(c echo.Context) error {
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

	id, _ := strconv.Atoi(c.Param("category_id"))

	fundraisings, err := h.fundraisingService.FindFundraisingByCategoryID(id, limit, offset)

	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraisings", err.Error()))
	}

	response := dto.ToAllFundraisingsResponse(fundraisings)

	return c.JSON(200, helper.ResponseWithData(true, "fundraisings retrieved successfully", response))
}
