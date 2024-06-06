package handler

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"
	"time"

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

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	organization := entities.Organization{
		Name:        req.Name,
		Avatar:      req.Avatar,
		Description: req.Description,
		IsVerified:  req.IsVerified,
		StartDate:   startDate,
		Contact:     req.Contact,
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
