package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ApplicationHandler struct {
	applicationService service.ApplicationService
}

func NewApplicationHandler(applicationService service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{applicationService: applicationService}
}

func (h *ApplicationHandler) RegisterApplication(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid id", err.Error()))
	}

	var request dto.ApplicationRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	igFile, err := c.FormFile("ig_image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "Invalid IG proof file", err.Error()))
	}

	ytFile, err := c.FormFile("yt_image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "Invalid YT proof file", err.Error()))
	}

	igURL, err := helper.UploadToCloudinary(igFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "Failed to upload IG proof", err.Error()))
	}

	ytURL, err := helper.UploadToCloudinary(ytFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "Failed to upload YT proof", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	application := request.ToEntity(igURL, ytURL, uint(userID), uint(id))

	createdApplication, err := h.applicationService.RegisterApplication(application)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create application", err.Error()))
	}

	response := dto.ToApplicationResponse(createdApplication)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "application created successfully", response))
}

func (h *ApplicationHandler) GetAllApplications(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 6
	}

	applications, total, err := h.applicationService.GetAllApplications(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get applications", err.Error()))
	}

	var responses []dto.ApplicationResponse
	for _, application := range applications {
		responses = append(responses, dto.ToApplicationResponse(application))
	}

	response := map[string]interface{}{
		"total": total,
		"data":  responses,
		"page":  page,
		"limit": limit,
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "applications retrieved successfully", response))
}

func (h *ApplicationHandler) GetApplicationByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid application ID", err.Error()))
	}

	application, err := h.applicationService.GetApplicationByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get application", err.Error()))
	}

	response := dto.ToApplicationResponse(application)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "application retrieved successfully", response))
}

func (h *ApplicationHandler) DeleteApplicationByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid application ID", err.Error()))
	}

	err = h.applicationService.DeleteApplicationByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete application", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "application deleted successfully"))
}
