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
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}
	userId, err := helper.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}
	var request dto.CommentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	comment := request.ToEntity(uint(articleId), uint(userId))

	createdComment, err := h.commentService.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to create comment", err.Error()))
	}

	// Remove CreatedAt, UpdatedAt, DeletedAt from the response
	response := map[string]interface{}{
		"id":          createdComment.ID,
		"customer_id": createdComment.UserID,
		"article_id":  createdComment.ArticleID,
		"comment":     createdComment.Comment,
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "comment created successfully", response))
}

func (h *CommentHandler) GetCommentsByArticleID(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 6 {
		limit = 6
	}

	comments, total, err := h.commentService.GetAllByArticleID(uint(articleId), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to get comments", err.Error()))
	}

	response := dto.ToCommentResponsesList(comments)
	totalInt64 := int64(total)
	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "article comments retrieved successfully", response, page, limit, totalInt64))
}

// func (h *CommentHandler) UpdateComment(c echo.Context) error {
// 	var request dto.CommentRequest
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid request", err.Error()))
// 	}

// 	if request.CustomerID == 0 {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "customer_id is required", "customer_id is missing"))
// 	}

// 	comment := request.ToEntity()

// 	updatedComment, err := h.commentService.UpdateComment(comment)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to update comment", err.Error()))
// 	}

// 	// Remove CreatedAt, UpdatedAt, DeletedAt from the response
// 	response := map[string]interface{}{
// 		"id":          updatedComment.ID,
// 		"customer_id": updatedComment.CustomerID,
// 		"article_id":  updatedComment.ArticleID,
// 		"comment":     updatedComment.Comment,
// 	}

// 	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "comment updated successfully", response))
// }

// func (h *CommentHandler) GetCommentByID(c echo.Context) error {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
// 	}

// 	comment, err := h.commentService.FindByID(uint(id))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse(false, "comment not found", err.Error()))
// 	}

// 	// Construct the response without unnecessary fields
// 	response := map[string]interface{}{
// 		"id":          comment.ID,
// 		"customer_id": comment.CustomerID,
// 		"article_id":  comment.ArticleID,
// 		"comment":     comment.Comment,
// 	}

// 	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "comment retrieved successfully", response))
// }

// func (h *CommentHandler) GetAllComments(c echo.Context) error {
// 	page, _ := strconv.Atoi(c.QueryParam("page"))
// 	limit, _ := strconv.Atoi(c.QueryParam("limit"))

// 	if page <= 0 {
// 		page = 1
// 	}
// 	if limit <= 0 || limit > 6 {
// 		limit = 6
// 	}

// 	comments, total, err := h.commentService.FindAll(page, limit)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to retrieve comments", err.Error()))
// 	}

// 	// Remove unnecessary fields from each comment
// 	var responseData []map[string]interface{}
// 	for _, comment := range comments {
// 		responseData = append(responseData, map[string]interface{}{
// 			"id":          comment.ID,
// 			"customer_id": comment.CustomerID,
// 			"article_id":  comment.ArticleID,
// 			"comment":     comment.Comment,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "comments retrieved successfully", responseData, page, limit, int64(total)))
// }

// func (h *CommentHandler) DeleteComment(c echo.Context) error {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid ID format", err.Error()))
// 	}

// 	err = h.commentService.DeleteComment(uint(id))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to delete comment", err.Error()))
// 	}

// 	return c.JSON(http.StatusOK, helper.GeneralResponse(true, "comment deleted successfully"))
// }
