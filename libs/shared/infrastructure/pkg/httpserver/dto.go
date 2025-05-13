package httpserver

type ErrorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status bool `json:"status"`
	Data   *any `json:"data"`
}

type ErrorResponse struct {
	SuccessResponse
	ErrorPayload
}
