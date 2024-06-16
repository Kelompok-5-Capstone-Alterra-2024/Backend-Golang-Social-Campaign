package handler

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	middleware "capstone/middlewares"
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
	articleService   service.ArticleService
	commentService   service.CommentService
}

func NewAdminHandler(adminService service.AdminService, volunteerService service.VolunteerService, articleService service.ArticleService, commentService service.CommentService) *AdminHandler {
	return &AdminHandler{adminService, volunteerService, articleService, commentService}
}

// func (h *AdminHandler) Login(c echo.Context) error {
// 	var request dto.LoginRequest
// 	c.Bind(&request)
// 	admin, err := h.adminService.Login(request)
// 	if err != nil {
// 		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid username or password"))
// 	}
// 	return c.JSON(200, helper.ResponseWithData(true, "Admin logged in successfully", admin.Token))
// }

func (h *AdminHandler) Login(c echo.Context) error {
	var request dto.LoginRequest
	c.Bind(&request)
	_, accessToken, refreshToken, err := h.adminService.Login(request)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "validation failed", "invalid credentials"))
	}

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	return c.JSON(200, helper.ResponseWithData(true, "Admin logged in successfully", response))
}

func (h *AdminHandler) RefreshTokenAdmin(c echo.Context) error {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	c.Bind(&request)

	claims, err := middleware.VerifyRefreshToken(request.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid refresh token",
		})
	}

	accessToken, refreshToken := middleware.GenerateToken(claims.ID, claims.Username, claims.Role)

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(200, helper.ResponseWithData(true, "Token refreshed successfully", response))
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

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	var req dto.CreateFundraisingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	imgFile, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid image url", err.Error()))
	}

	imageUrl, err := helper.UploadToCloudinary(imgFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to upload image", err.Error()))
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
		ImageUrl:              imageUrl,
		Title:                 req.Title,
		GoalAmount:            req.TargetAmount,
		Description:           req.Description,
		StartDate:             startDate,
		EndDate:               endDate,
		FundraisingCategoryID: req.CategoryID,
		OrganizationID:        req.OrganizationID,
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

func (h *AdminHandler) GetOrganizationByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	organization, err := h.adminService.GetOrganizationByID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get organization", err.Error()))
	}

	response := dto.ToAdminOrgResponse(organization)

	return c.JSON(200, helper.ResponseWithData(true, "organization retrieved successfully", response))
}

func (h *AdminHandler) GetFundraisingByOrganization(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}
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
	fundraisings, err := h.adminService.GetFundraisingByOrganizationID(id, limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraising", err.Error()))
	}

	response := dto.ToAdminAllOrgFundraisingResponse(fundraisings)

	return c.JSON(200, helper.ResponseWithPagination("success", "fundraising retrieved successfully", response, page, limit, int64(len(fundraisings))))
}

func (h *AdminHandler) GetVolunteerByOrganization(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}
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

	volunteers, err := h.adminService.GetVolunteerByOrganizationID(id, limit, offset)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get fundraising", err.Error()))
	}

	response := dto.ToAdminAllOrgVolunteersResponse(volunteers)

	return c.JSON(200, helper.ResponseWithPagination("success", "volunteer retrieved successfully", response, page, limit, int64(len(volunteers))))
}

func (h *AdminHandler) EditOrganization(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.OrganizationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	imgFile, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid image url", err.Error()))
	}

	imageUrl, err := helper.UploadToCloudinary(imgFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to upload image", err.Error()))
	}

	startDate, err := time.Parse("2006-01-02", req.JoinDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid start date format")
	}

	organization := entities.Organization{
		Name:        req.Name,
		Description: req.Description,
		StartDate:   startDate,
		IsVerified:  req.IsVerified,
		Website:     req.Website,
		Instagram:   req.Instagram,
		NoRekening:  req.NoRekening,
		Avatar:      imageUrl,
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

func (h *AdminHandler) GetUserDonations(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	userDonations, total, err := h.adminService.GetDonationsByUserID(id, page, limit)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get user donations", err.Error()))
	}

	response := dto.ToAdminAllUserDonationResponse(userDonations)
	return c.JSON(200, helper.ResponseWithPagination("success", "user donations retrieved successfully", response, page, limit, int64(total)))
}

func (h *AdminHandler) GetUserVolunteers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	userVolunteers, total, err := h.adminService.GetVolunteersByUserID(id, page, limit)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get user volunteers", err.Error()))
	}

	response := dto.ToAdminAllUserVolunteers(userVolunteers)
	return c.JSON(200, helper.ResponseWithPagination("success", "user volunteers retrieved successfully", response, page, limit, int64(total)))
}

func (h *AdminHandler) GetUserDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}

	user, err := h.adminService.GetUserByID(id)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get user", err.Error()))
	}

	response := dto.ToAdminUserDetailResponse(user)
	return c.JSON(200, helper.ResponseWithData(true, "user retrieved successfully", response))
}

func (h *AdminHandler) EditUsers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid user id", err.Error()))
	}

	var req dto.EditUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	user := req.ToEntity()
	user.ID = uint(id)

	_, err = h.adminService.UpdateUserByID(uint(id), user)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to edit user", err.Error()))
	}

	return c.JSON(200, helper.GeneralResponse(true, "user edited successfully"))

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

func (h *AdminHandler) GetAdminAllVolunteers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 10
	}
	volunteers, total, err := h.volunteerService.FindAll(page, limit)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get volunteers", err.Error()))
	}

	response := dto.ToAdminAllVolunteersResponse(volunteers)
	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "volunteers retrieved successfully", response, page, limit, int64(total)))
}

func (h *AdminHandler) GetAdminAllArticle(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 10
	}

	articles, total, err := h.articleService.FindAll(page, limit)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get articles", err.Error()))
	}

	comments, err := h.commentService.GetAllComments()
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get comments", err.Error()))
	}

	response := dto.ToAdminAllArticleResponses(articles, comments)
	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "articles retrieved successfully", response, page, limit, int64(total)))
}

func (h *AdminHandler) GetAdminArticleByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	article, err := h.articleService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	comments, err := h.commentService.GetAllComments()
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get comments", err.Error()))
	}

	response := dto.ToAdminArticleResponses(article, comments)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "article retrieved successfully", response))
}

func (h *AdminHandler) GetAllDonationManual(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 10
	}

	donations, total, err := h.adminService.GetAllDonations(page, limit)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donations", err.Error()))
	}

	response := dto.ToAdminAllDonationResponses(donations)
	totalInt64 := int64(total)
	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "donations retrieved successfully", response, page, limit, totalInt64))
}

func (h *AdminHandler) InputAmountDonationManual(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	type InputAmount struct {
		Amount int `json:"amount"`
	}

	var input InputAmount
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	_, err = h.adminService.AddAmountToUserDonation(uint(id), input.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "donation amount added successfully"))
}

func (h *AdminHandler) GetDailyDonationSummary(c echo.Context) error {

	donations, err := h.adminService.GetDailyDonationSummary()
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get donations", err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "donations retrieved successfully", donations))
}

func (h *AdminHandler) GetDataTotalContent(c echo.Context) error {

	data, err := h.adminService.GetDataTotalContent()
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get data", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "data retrieved successfully", data))
}

func (h *AdminHandler) GetArticlesOrderedByBookmarks(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	articles, bookmarkCounts, _, err := h.adminService.GetArticlesOrderedByBookmarks(pageInt, limitInt)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get articles", err.Error()))
	}

	// Create a custom response format
	response := make([]map[string]interface{}, len(articles))
	for i, article := range articles {
		response[i] = map[string]interface{}{
			"image":          article.ImageURL,
			"title":          article.Title,
			"content":        article.Content,
			"bookmark_count": bookmarkCounts[i],
		}
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "articles retrieved successfully", response))

}
