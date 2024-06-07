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
	volunteerService   service.VolunteerService
	applicationService service.ApplicationService
}

func NewVolunteerHandler(volunteerService service.VolunteerService, applicationService service.ApplicationService) *VolunteerHandler {
	return &VolunteerHandler{volunteerService: volunteerService, applicationService: applicationService}
}

func (h *VolunteerHandler) CreateVolunteer(c echo.Context) error {
	var request dto.VolunteerRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	imgFile, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid image url", err.Error()))
	}

	imageUrl, err := helper.UploadToCloudinary(imgFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to upload image", err.Error()))
	}

	volunteer, err := request.ToEntity(imageUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid date format", err.Error()))
	}

	_, err = h.volunteerService.CreateVolunteer(volunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create volunteer", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "volunteer created successfully"))
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

	applications, err := h.applicationService.GetApplicationByVacancyID(uint(id))

	response := dto.ToVolunteerResponse(volunteer, applications)
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

	response := dto.ToAdminAllVolunteersResponse(volunteers)

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "volunteers retrieved successfully", response, page, limit, int64(total)))
}

func (h *VolunteerHandler) ApplyForVolunteer(c echo.Context) error {
	volunteerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid volunteer ID format", err.Error()))
	}

	userId, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	_, err = h.volunteerService.ApplyForVolunteer(uint(volunteerID), uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to apply for volunteer", err.Error()))
	}

	// response := dto.ToVolunteerResponse(updatedVolunteer)
	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "volunteer application successful"))
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

	imgFile, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid image url", err.Error()))
	}

	imageUrl, err := helper.UploadToCloudinary(imgFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to upload image", err.Error()))
	}

	volunteer, err := request.ToEntity(imageUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid date format", err.Error()))
	}

	volunteer.ID = uint(id)
	_, err = h.volunteerService.UpdateVolunteer(volunteer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to update volunteer", err.Error()))
	}

	// response := dto.ToVolunteerResponse(updatedVolunteer)
	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "volunteer updated successfully"))
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

func (h *VolunteerHandler) ConfirmVolunteer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid id", err.Error()))
	}

	confirmVolunteer, err := h.volunteerService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "volunteer not found", err.Error()))
	}

	userId, err := helper.GetUserIDFromJWT(c)

	response := dto.ToConfirmVolunteerResponse(confirmVolunteer, uint(userId))
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "volunteer confirmed successfully", response))
}
