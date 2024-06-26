package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	var request dto.ArticleRequest
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

	article := request.ToEntity(imageUrl)

	_, err = h.articleService.CreateArticle(article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	// response := dto.ToArticleResponse(createdArticle)
	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "article created successfully"))
}

func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	var request dto.ArticleRequest
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

	article := request.ToEntity(imageUrl)
	article.ID = uint(id)

	_, err = h.articleService.UpdateArticle(article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	// response := dto.ToArticleResponse(updatedArticle)
	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "article updated successfully"))
}

func (h *ArticleHandler) GetArticleByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	article, err := h.articleService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	response := dto.ToArticleResponses(article)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "article retrieved successfully", response))
}

func (h *ArticleHandler) GetAllArticles(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	articles, total, err := h.articleService.FindAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	response := dto.ToArticleResponsesList(articles)
	totalInt64 := int64(total)
	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "articles retrieved successfully", response, page, limit, totalInt64))
}

func (h *ArticleHandler) DeleteArticle(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	err = h.articleService.DeleteArticle(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "article deleted successfully"))
}

func (h *ArticleHandler) GetTopArticles(c echo.Context) error {
	articles, err := h.articleService.GetTopArticles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	response := dto.ToArticleResponsesList(articles)
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "top articles retrieved successfully", response))
}
