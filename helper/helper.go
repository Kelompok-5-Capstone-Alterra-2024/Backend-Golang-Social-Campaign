package helper

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type generalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GeneralResponse(status, message string) generalResponse {
	messageRes := generalResponse{
		Status:  status,
		Message: message,
	}
	return messageRes
}

type responseWithData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(status, message string, data interface{}) responseWithData {
	messageRes := responseWithData{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return messageRes
}

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrorResponse(status, message string, err any) errorResponse {
	messageRes := errorResponse{
		Status:  status,
		Message: message,
		Error:   err,
	}

	return messageRes
}

func GetToken(auth string) string {
	splittedAuth := strings.Split(auth, "Bearer ")
	return splittedAuth[1]
}

func DecodePayload(token string) (map[string]interface{}, error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT token format")
	}

	// Decode base64 encoded payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var payloadMap map[string]interface{}
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		return nil, err
	}

	return payloadMap, nil
}
