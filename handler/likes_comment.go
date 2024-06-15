package handler

import (
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
	// var request dto.LikesCommentRequest
	// if err := c.Bind(&request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	// }

	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}
	// request.CommentID = uint(commentID)

	// likesComment := request.ToEntity(uint(commentID), uint(userID))

	// createdLikesComment, err := h.likesCommentService.CreateLikesComment(c.Request().Context(), uint(commentID), uint(userID))
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create like on comment", err.Error()))
	// }

	// response := dto.LikesCommentResponse{
	// 	ID:        createdLikesComment.ID,
	// 	UserID:    createdLikesComment.UserID,
	// 	CommentID: createdLikesComment.CommentID,
	// }

	if err := h.likesCommentService.LikesComment(c.Request().Context(), uint(commentID), uint(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create like on comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "like on comment created successfully"))
}

func (h *LikesCommentHandler) DeleteLikesComment(c echo.Context) error {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userID, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	}

	if err := h.likesCommentService.DeleteLikesComment(c.Request().Context(), uint(commentID), uint(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete like on comment", err.Error()))
	}

	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "like on comment deleted successfully"))
}

// func (h *LikesCommentHandler) GetLikesCommentByID(c echo.Context) error {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
// 	}

// 	like, err := h.likesCommentService.GetLikesCommentByID(uint(id))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to fetch like on comment", err.Error()))
// 	}

// 	response := struct {
// 		Success bool                       `json:"success"`
// 		Message string                     `json:"message"`
// 		Data    []dto.LikesCommentResponse `json:"data"`
// 	}{
// 		Success: true,
// 		Message: "like on comment retrieved successfully",
// 		Data: []dto.LikesCommentResponse{
// 			{
// 				ID:         like.ID,
// 				CustomerID: like.UserID,
// 				CommentID:  like.CommentID,
// 			},
// 		},
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func (h *LikesCommentHandler) GetAllLikesComments(c echo.Context) error {
// 	likes, err := h.likesCommentService.GetAllLikesComments()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to fetch likes on comments", err.Error()))
// 	}

// 	response := struct {
// 		Success bool                       `json:"success"`
// 		Message string                     `json:"message"`
// 		Data    []dto.LikesCommentResponse `json:"data"`
// 	}{
// 		Success: true,
// 		Message: "likes on comments retrieved successfully",
// 		Data:    make([]dto.LikesCommentResponse, len(likes)),
// 	}

// 	for i, like := range likes {
// 		response.Data[i] = dto.LikesCommentResponse{
// 			ID:         like.ID,
// 			CustomerID: like.UserID,
// 			CommentID:  like.CommentID,
// 		}
// 	}

// 	return c.JSON(http.StatusOK, response)
// }
