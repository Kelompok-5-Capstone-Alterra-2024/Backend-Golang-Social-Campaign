package helper

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
