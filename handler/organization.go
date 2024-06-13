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

type OrganizationHandler struct {
	organizationService service.OrganizationService
}

func NewOrganizationHandler(organizationService service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{organizationService}
}

func (h *OrganizationHandler) CreateOrganization(c echo.Context) error {
	var req dto.OrganizationRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create organization", err.Error()))
	}

	fileHeader, _ := c.FormFile("avatar")
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := "cloudinary://633714464826515:u1W6hqq-Gb8y-SMpXe7tzs4mH44@dvrhf8d9t"
	cloudinaryUsecase, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cloudinaryUsecase.Upload.Upload(ctx, file, uploader.UploadParams{})

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	organization := entities.Organization{
		Name:        req.Name,
		Avatar:      response.SecureURL,
		Description: req.Description,
		IsVerified:  req.IsVerified,
		StartDate:   startDate,
		Website:     req.Website,
		Instagram:   req.Instagram,
		NoRekening:  req.NoRekening,
	}

	_, err = h.organizationService.CreateOrganization(organization)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create organization", err.Error()))
	}

	return c.JSON(201, helper.GeneralResponse(true, "organization created successfully"))
}

func (h *OrganizationHandler) GetOrganizations(c echo.Context) error {
	organizations, err := h.organizationService.FindOrganizations()
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get organizations", err.Error()))
	}
	return c.JSON(200, helper.ResponseWithData(true, "organizations retrieved successfully", organizations))
}

func (h *OrganizationHandler) GetOrganizationByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	organization, err := h.organizationService.FindOrganizationByID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get organization", err.Error()))
	}

	fundraisings, err := h.organizationService.FindFundraisingByOrganizationID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get organization", err.Error()))
	}

	response := dto.ToOrganizationResponse(organization, fundraisings)
	return c.JSON(200, helper.ResponseWithData(true, "organization retrieved successfully", response))
}
