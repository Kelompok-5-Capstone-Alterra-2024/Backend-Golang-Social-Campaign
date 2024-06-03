package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VolunteerHandler struct {
	volunteerService service.VolunteerService
}

func NewVolunteerHandler(volunteerService service.VolunteerService) *VolunteerHandler {
	return &VolunteerHandler{volunteerService: volunteerService}
}

func (h *VolunteerHandler) CreateVolunteer(c echo.Context) error {
	var request dto.VolunteerRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	volunteer, err := request.ToEntity()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid date format", err.Error()))
	}

	createdVolunteer, err := h.volunteerService.CreateVolunteer(volunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer created successfully", createdVolunteer))
}

func (h *VolunteerHandler) GetVolunteerByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	volunteer, err := h.volunteerService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "volunteer not found", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer retrieved successfully", volunteer))
}

func (h *VolunteerHandler) GetAllVolunteers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	volunteers, total, err := h.volunteerService.FindAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to retrieve volunteers", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "volunteers retrieved successfully", volunteers, page, limit, total))
}
