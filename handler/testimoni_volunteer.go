package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TestimoniVolunteerHandler struct {
	testimoniVolunteerService service.TestimoniVolunteerService
}

func NewTestimoniVolunteerHandler(testimoniVolunteerService service.TestimoniVolunteerService) *TestimoniVolunteerHandler {
	return &TestimoniVolunteerHandler{testimoniVolunteerService: testimoniVolunteerService}
}

func (h *TestimoniVolunteerHandler) CreateTestimoniVolunteer(c echo.Context) error {
	var request dto.TestimoniVolunteerRequest
	volunteer_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid volunteer ID ", err.Error()))
	}

	userId, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	// Add validation here
	customerJoined := h.testimoniVolunteerService.CustomerJoinedVolunteer(uint(userId), uint(volunteer_id))
	if !customerJoined {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "customer has not joined the volunteer", ""))
	}

	alreadyTestified := h.testimoniVolunteerService.HasCustomerGivenTestimony(uint(userId), uint(volunteer_id))
	if alreadyTestified {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "customer has already given a testimony", ""))
	}

	testimoniVolunteer := request.ToEntity(uint(volunteer_id), uint(userId))

	createdTestimoniVolunteer, err := h.testimoniVolunteerService.CreateTestimoniVolunteer(testimoniVolunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create testimoni volunteer", err.Error()))
	}

	// Convert created_at to Asia/Jakarta timezone
	loc, _ := time.LoadLocation("Asia/Jakarta")
	createdAt := createdTestimoniVolunteer.CreatedAt.In(loc).Format("02/01/2006 15:04:05 MST")

	response := map[string]interface{}{
		"user_id":      createdTestimoniVolunteer.UserID,
		"volunteer_id": createdTestimoniVolunteer.VacancyID,
		"testimoni":    createdTestimoniVolunteer.Testimoni,
		"rating":       createdTestimoniVolunteer.Rating,
		"date":         createdAt,
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "testimoni volunteer created successfully", response))
}

func (h *TestimoniVolunteerHandler) GetTestimoniVolunteerByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID", err.Error()))
	}

	testimoniVolunteer, err := h.testimoniVolunteerService.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "testimoni volunteer not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "error retrieving testimoni volunteer", err.Error()))
	}

	response := dto.ToTestimoniVolunteerResponse(testimoniVolunteer)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "retrieved testimoni volunteer successfully", response))
}

func (h *TestimoniVolunteerHandler) GetAllTestimoniVolunteers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 6
	}

	testimoniVolunteers, total, err := h.testimoniVolunteerService.FindAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "error retrieving testimoni volunteers", err.Error()))
	}

	responses := make([]dto.TestimoniVolunteerResponse, len(testimoniVolunteers))
	for i, tv := range testimoniVolunteers {
		responses[i] = dto.ToTestimoniVolunteerResponse(tv)
	}

	response := map[string]interface{}{
		"total": total,
		"data":  responses,
		"page":  page,
		"limit": limit,
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "retrieved all testimoni volunteers successfully", response))
}

func (h *TestimoniVolunteerHandler) DeleteTestimoniVolunteer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID", err.Error()))
	}

	err = h.testimoniVolunteerService.DeleteTestimoniVolunteer(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "testimoni volunteer not found", ""))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete testimoni volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "testimoni volunteer deleted successfully"))
}
