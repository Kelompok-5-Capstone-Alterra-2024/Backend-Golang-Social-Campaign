package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TestimoniVolunteerHandler struct {
	testimoniVolunteerService service.TestimoniVolunteerService
}

func NewTestimoniVolunteerHandler(testimoniVolunteerService service.TestimoniVolunteerService) *TestimoniVolunteerHandler {
	return &TestimoniVolunteerHandler{testimoniVolunteerService: testimoniVolunteerService}
}

func (h *TestimoniVolunteerHandler) CreateTestimoniVolunteer(c echo.Context) error {
	var request dto.TestimoniVolunteerRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	testimoniVolunteer, err := request.ToEntity()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid data format", err.Error()))
	}

	createdTestimoniVolunteer, err := h.testimoniVolunteerService.CreateTestimoniVolunteer(testimoniVolunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create testimoni volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "testimoni volunteer created successfully", createdTestimoniVolunteer))
}

func (h *TestimoniVolunteerHandler) GetTestimoniVolunteerByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	testimoniVolunteer, err := h.testimoniVolunteerService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "testimoni volunteer not found", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "testimoni volunteer retrieved successfully", testimoniVolunteer))
}

func (h *TestimoniVolunteerHandler) GetAllTestimoniVolunteers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	testimoniVolunteers, total, err := h.testimoniVolunteerService.FindAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to retrieve testimoni volunteers", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "testimoni volunteers retrieved successfully", testimoniVolunteers, page, limit, int64(total)))
}

func (h *TestimoniVolunteerHandler) DeleteTestimoniVolunteer(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	err = h.testimoniVolunteerService.DeleteTestimoniVolunteer(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete testimoni volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "testimoni volunteer deleted successfully"))
}
