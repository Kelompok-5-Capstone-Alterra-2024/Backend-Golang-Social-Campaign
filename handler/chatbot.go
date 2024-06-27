package handler

import (
	"capstone/dto"
	"capstone/helper"
	"capstone/service"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatBotHandler struct {
	chatBotService service.ChatbotService
}

func NewChatBotHandler(chatBotService service.ChatbotService) *ChatBotHandler {
	return &ChatBotHandler{chatBotService: chatBotService}
}

func (h *ChatBotHandler) CreateChatBot(c echo.Context) error {
	var request dto.ChatBotRequest
	c.Bind(&request)

	chatBot := dto.Chatbot{
		ChatID:  request.ID,
		Message: request.Message,
	}

	res, err := h.chatBotService.Create(chatBot)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to create chatbot", err.Error()))
	}

	var response dto.ChatbotResponse

	response.ID = res.ID
	response.ChatID = res.ChatID
	response.Role = res.Role
	response.Message = res.Message
	response.CreatedAt = res.CreatedAt.Format(time.DateTime)

	return c.JSON(200, helper.ResponseWithData(true, "chatbot created successfully", response))
}

func (h *ChatBotHandler) GetChatBot(c echo.Context) error {
	// userId, err := helper.GetUserIDFromJWT(c)
	// if err != nil {
	// 	return c.JSON(401, helper.ErrorResponse(false, "unauthorized", err.Error()))
	// }

	chatId := c.Param("chat_id")

	chatBot, err := h.chatBotService.GetByID(chatId)
	if err != nil {
		return c.JSON(500, helper.ErrorResponse(false, "failed to get chatbot", err.Error()))
	}

	var response []dto.ChatbotResponse
	for _, v := range chatBot {
		var res dto.ChatbotResponse
		res.ID = v.ID
		res.ChatID = v.ChatID
		res.Role = v.Role
		res.Message = v.Message
		res.CreatedAt = v.CreatedAt.Format(time.DateTime)
		response = append(response, res)
	}

	return c.JSON(200, helper.ResponseWithData(true, "chatbot retrieved successfully", response))
}
