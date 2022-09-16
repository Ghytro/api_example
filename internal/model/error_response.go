package model

type ErrorResponse struct {
	OriginErr  error `json:"-"`
	StatusCode int   `json:"-"`

	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.OriginErr.Error()
}

func GenErrorResponse(err error, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		OriginErr:  err,
		Message:    err.Error(),
	}
}
