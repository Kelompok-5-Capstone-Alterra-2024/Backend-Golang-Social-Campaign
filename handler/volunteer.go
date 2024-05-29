package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"

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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed", "invalid request", err.Error()))
	}

	volunteer := request.ToEntity()
	createdVolunteer, err := h.volunteerService.CreateVolunteer(volunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed", "failed to create volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.NewDataResponse("success", createdVolunteer))
}

func (h *VolunteerHandler) GetVolunteerByID(c echo.Context) error {
	id, err := helper.StringToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed", "invalid request", err.Error()))
	}

	volunteer, err := h.volunteerService.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed", "failed to get volunteer", err.Error()))
	}
	return c.JSON(http.StatusOK, helper.NewDataResponse("success", volunteer))
}

func (h *VolunteerHandler) GetAllVolunteer(c echo.Context) error {
	volunteers, err := h.volunteerService.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed", "failed to get volunteers", err.Error()))
	}
	return c.JSON(http.StatusOK, helper.NewDataResponse("success", volunteers))
}
