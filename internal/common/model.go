package common

type errResponse struct {
	TraceID string `json:"trace_id"`
}

type response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
	Error   *errResponse `json:"error,omitempty"`
}

func NewSuccessResponse(status int, message string, data interface{}) *response {
	return &response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(status int, message string, traceID string) *response {
	return &response{
		Status:  status,
		Message: message,
		Error: &errResponse{
			TraceID: traceID,
		},
	}
}
