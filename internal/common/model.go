package common

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func NewSuccessResponse(status int, message string, data interface{}) *response {
	return &response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(status int, message string, err string) *response {
	return &response{
		Status:  status,
		Message: message,
		Error:   err,
	}
}
