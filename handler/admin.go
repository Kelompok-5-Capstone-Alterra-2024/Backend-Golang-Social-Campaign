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
	adminService     service.AdminService
	volunteerService service.VolunteerService
}

func NewAdminHandler(adminService service.AdminService, volunteerService service.VolunteerService) *AdminHandler {
	return &AdminHandler{adminService, volunteerService}
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

	fileHeader, _ := c.FormFile("image_url")
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := "cloudinary://633714464826515:u1W6hqq-Gb8y-SMpXe7tzs4mH44@dvrhf8d9t"
	cloudinaryUsecase, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cloudinaryUsecase.Upload.Upload(ctx, file, uploader.UploadParams{})

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid end date format")
	}

	fundraising := entities.Fundraising{
		ImageUrl:              response.SecureURL,
		Title:                 req.Title,
		GoalAmount:            req.TargetAmount,
		Description:           req.Description,
		StartDate:             startDate,
		EndDate:               endDate,
		FundraisingCategoryID: req.CategoryID,
		OrganizationID:        req.OrganizationID,
		Status:                "unachieved",
	}

	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create fundraising", err.Error()))
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

	fileHeader, _ := c.FormFile("image_url")
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := "cloudinary://633714464826515:u1W6hqq-Gb8y-SMpXe7tzs4mH44@dvrhf8d9t"
	cloudinaryUsecase, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cloudinaryUsecase.Upload.Upload(ctx, file, uploader.UploadParams{})

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid end date format")
	}

	fundraising := entities.Fundraising{
		ImageUrl:              response.SecureURL,
		Title:                 req.Title,
		GoalAmount:            req.TargetAmount,
		Description:           req.Description,
		StartDate:             startDate,
		EndDate:               endDate,
		FundraisingCategoryID: req.CategoryID,
		OrganizationID:        req.OrganizationID,
	}

	_, err = h.adminService.SaveImageFundraising(uint(id), response.SecureURL)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to edit fundraising", err.Error()))
	}

	_, err = h.adminService.UpdateFundraising(uint(id), fundraising)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to edit fundraising", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "fundraising edited successfully"))
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

func (h *AdminHandler) GetAllUsers(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit
	users, err := h.adminService.GetUsers(limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get users", err.Error()))
	}

	response := dto.ToAdminAllUsersResponses(users)
	return c.JSON(200, helper.ResponseWithData(true, "users retrieved successfully", response))

}

func (h *AdminHandler) GetDetailUserWithDonations(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}

	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	userDonations, err := h.adminService.GetDonationsByUserID(id, limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get user donations", err.Error()))
	}

	return c.JSON(200, helper.ResponseWithData(true, "user donations retrieved successfully", userDonations))

}

func (h *AdminHandler) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}

	err = h.adminService.DeleteUserWithDonations(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete user", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "user deleted successfully"))
}

func (h *AdminHandler) GetAllVolunteers(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit
	volunteers, total, err := h.volunteerService.FindAll(limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get volunteers", err.Error()))
	}

	response := dto.ToAdminAllVolunteersResponse(volunteers)
	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "volunteers retrieved successfully", response, page, limit, int64(total)))
}
