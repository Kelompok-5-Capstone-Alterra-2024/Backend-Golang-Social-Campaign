package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikesCommentHandler struct {
	likesCommentService service.LikesCommentService
}

func NewLikesCommentHandler(likesCommentService service.LikesCommentService) *LikesCommentHandler {
	return &LikesCommentHandler{likesCommentService: likesCommentService}
}

func (h *LikesCommentHandler) CreateLikesComment(c echo.Context) error {
	var request dto.LikesCommentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed", "invalid request", err.Error()))
	}

	likesComment := request.ToEntity()

	createdLikesComment, err := h.likesCommentService.CreateLikesComment(likesComment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed", "failed to create like on comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData("success", "like on comment created successfully", createdLikesComment))
}

func (h *LikesCommentHandler) DeleteLikesComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed", "invalid ID format", err.Error()))
	}

	err = h.likesCommentService.DeleteLikesComment(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed", "failed to delete like on comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse("success", "like on comment deleted successfully"))
}
