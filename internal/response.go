package internal

func NewResponseOK(data interface{}) *response {
	return &response{
		Status:  200,
		Message: "ok",
		Data:    data,
		Errors:  []string{},
	}
}
