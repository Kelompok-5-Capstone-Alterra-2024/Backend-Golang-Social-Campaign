package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	var request dto.CommentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	comment := request.ToEntity()

	createdComment, err := h.commentService.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "comment created successfully", createdComment))
}

func (h *CommentHandler) UpdateComment(c echo.Context) error {
	var request dto.CommentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	comment := request.ToEntity()

	updatedComment, err := h.commentService.UpdateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to update comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "comment updated successfully", updatedComment))
}

func (h *CommentHandler) GetCommentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	comment, err := h.commentService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "comment not found", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "comment retrieved successfully", comment))
}

func (h *CommentHandler) GetAllComments(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	comments, total, err := h.commentService.FindAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to retrieve comments", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "comments retrieved successfully", comments, page, limit, int64(total)))
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	err = h.commentService.DeleteComment(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "comment deleted successfully"))
}
