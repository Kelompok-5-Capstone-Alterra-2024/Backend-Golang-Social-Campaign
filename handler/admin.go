package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{adminService}
}

func (h *AdminHandler) Login(c echo.Context) error {
	var request dto.LoginRequest
	c.Bind(&request)
	admin, err := h.adminService.Login(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid credentials"))
	}
	return c.JSON(200, helper.ResponseWithData(true, "Admin logged in successfully", admin.Token))
}
