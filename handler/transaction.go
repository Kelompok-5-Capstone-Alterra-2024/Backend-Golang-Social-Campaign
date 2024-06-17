package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {

	fundraisingID, _ := strconv.Atoi(c.Param("id"))

	var request dto.DistributeFundFundraisingRequest
	c.Bind(&request)

	imgFile, err := c.FormFile("image_payment")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "invalid image url", err.Error()))
	}

	imageUrl, err := helper.UploadToCloudinary(imgFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(false, "failed to upload image", err.Error()))
	}

	request.FundraisingID = uint(fundraisingID)
	request.ImagePayment = imageUrl

	transaction, err := h.transactionService.CreateTransaction(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "failed to create transaction", err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.ResponseWithData(true, "Transaction created", transaction))
}

func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.transactionService.GetTransactionByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "failed to get transaction", err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseWithData(true, "Transaction found", transaction))
}

func (h *TransactionHandler) GetTransactions(c echo.Context) error {
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

	transactions, err := h.transactionService.GetTransactions(limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(false, "failed to get transactions", err.Error()))
	}

	response := dto.ToTransactionHistoriesRespone(transactions)

	return c.JSON(http.StatusOK, helper.ResponseWithPagination("success", "transactions retrieved successfully", response, page, limit, int64(len(transactions))))
}
