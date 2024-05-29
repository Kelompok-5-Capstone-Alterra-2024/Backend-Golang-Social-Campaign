package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApplicationHandler struct {
	applicationService service.ApplicationService
}

func NewApplicationHandler(applicationService service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{applicationService}
}

func (h *ApplicationHandler) RegisterApplication(c echo.Context) error {
	var request dto.ApplicationRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed", "invalid request", err.Error()))
	}

	application := request.ToEntity()
	_, err := h.applicationService.RegisterApplication(application)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed", "registration failed", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse("success", "Application registered successfully"))
}
