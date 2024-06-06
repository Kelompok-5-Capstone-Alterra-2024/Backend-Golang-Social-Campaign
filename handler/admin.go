package handler

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	"capstone/service"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid username or password"))
	}
	return c.JSON(200, helper.ResponseWithData(true, "Admin logged in successfully", admin.Token))
}

func (h *AdminHandler) GetFundraisings(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 8
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	fundraisings, err := h.adminService.GetFundraisings(limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraisings", err.Error()))
	}
	response := dto.ToAdminAllFundraisingsResponse(fundraisings)
	return c.JSON(200, helper.ResponseWithData(true, "fundraisings retrieved successfully", response))
}

func (h *AdminHandler) CreateFundraisingContent(c echo.Context) error {

	var req dto.CreateFundraisingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid end date format")
	}

	fundraising := entities.Fundraising{
		ImageUrl:              req.ImageUrl,
		Title:                 req.Title,
		GoalAmount:            req.TargetAmount,
		Description:           req.Description,
		StartDate:             startDate,
		EndDate:               endDate,
		FundraisingCategoryID: req.CategoryID,
		OrganizationID:        req.OrganizationID,
		Status:                "unachieved",
	}

	_, err = h.adminService.CreateFudraising(c.Request().Context(), fundraising)

	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create fundraising", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "fundraising created successfully"))

}

func (h *AdminHandler) DeleteFundraising(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	err := h.adminService.DeleteFundraising(uint(id))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to delete fundraising", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "fundraising deleted successfully"))
}

func (h *AdminHandler) EditFundraising(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.CreateFundraisingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid end date format")
	}

	fundraising := entities.Fundraising{
		ImageUrl:              req.ImageUrl,
		Title:                 req.Title,
		GoalAmount:            req.TargetAmount,
		Description:           req.Description,
		StartDate:             startDate,
		EndDate:               endDate,
		FundraisingCategoryID: req.CategoryID,
		OrganizationID:        req.OrganizationID,
	}

	updatedFundraising, err := h.adminService.UpdateFundraising(uint(id), fundraising)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to edit fundraising", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData(true, "fundraising edited successfully", updatedFundraising))
}

func (h *AdminHandler) GetDetailFundraising(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fundraising, err := h.adminService.GetFundraisingByID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraising", err.Error()))
	}
	response := dto.ToAdminFundraisingResponse(fundraising)
	return c.JSON(200, helper.ResponseWithData(true, "fundraising retrieved successfully", response))
}

func (h *AdminHandler) GetDonationsByFundraisingID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

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
	donations, err := h.adminService.GetDonationByFundraisingID(id, limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donation", err.Error()))
	}

	response := dto.ToAdminAllFundraisingDonationResponse(donations)
	return c.JSON(200, helper.ResponseWithData(true, "donation retrieved successfully", response))
}

func (h *AdminHandler) GetAllOrganizations(c echo.Context) error {

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
	organizations, err := h.adminService.GetOrganizations(limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get organization", err.Error()))
	}

	response := dto.ToAdminAllOrganizationsResponse(organizations)
	return c.JSON(200, helper.ResponseWithData(true, "organization retrieved successfully", response))
}

func (h *AdminHandler) EditOrganization(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.OrganizationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	fileHeader, _ := c.FormFile("avatar")
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := "cloudinary://633714464826515:u1W6hqq-Gb8y-SMpXe7tzs4mH44@dvrhf8d9t"
	cloudinaryUsecase, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cloudinaryUsecase.Upload.Upload(ctx, file, uploader.UploadParams{})

	organization := entities.Organization{
		Name:        req.Name,
		Description: req.Description,
		IsVerified:  req.IsVerified,
		Contact:     req.Contact,
		Avatar:      response.SecureURL,
	}

	_, err := h.adminService.SaveImageOraganization(uint(id), response.SecureURL)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to edit organization", err.Error()))
	}

	_, err = h.adminService.UpdateOrganization(uint(id), organization)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to edit organization", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "organization edited successfully"))
}

func (h *AdminHandler) DeleteOrganization(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.adminService.DeleteOrganization(uint(id))
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to delete organization", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "organization deleted successfully"))

}
