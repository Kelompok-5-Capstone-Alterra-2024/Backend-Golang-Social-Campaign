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

	response := dto.ToVolunteerResponse(createdVolunteer)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer created successfully", response))
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

	response := dto.ToVolunteerResponse(volunteer)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer retrieved successfully", response))
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

	var response []dto.VolunteerResponse
	for _, volunteer := range volunteers {
		response = append(response, dto.ToVolunteerResponse(volunteer))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "volunteers retrieved successfully", response, page, limit, int64(total)))
}

func (h *VolunteerHandler) ApplyForVolunteer(c echo.Context) error {
	volunteerID, err := strconv.ParseUint(c.Param("volunteer_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid volunteer ID format", err.Error()))
	}

	customerID, err := strconv.ParseUint(c.Param("customer_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid customer ID format", err.Error()))
	}

	updatedVolunteer, err := h.volunteerService.ApplyForVolunteer(uint(volunteerID), uint(customerID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to apply for volunteer", err.Error()))
	}

	response := dto.ToVolunteerResponse(updatedVolunteer)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer application successful", response))
}

func (h *VolunteerHandler) UpdateVolunteer(c echo.Context) error {
	var request dto.VolunteerRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	volunteer, err := request.ToEntity()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid date format", err.Error()))
	}

	volunteer.ID = uint(id)
	updatedVolunteer, err := h.volunteerService.UpdateVolunteer(volunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to update volunteer", err.Error()))
	}

	response := dto.ToVolunteerResponse(updatedVolunteer)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer updated successfully", response))
}

func (h *VolunteerHandler) DeleteVolunteer(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	err = h.volunteerService.DeleteVolunteer(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "volunteer deleted successfully"))
}
