package http_response


type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(code int, message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    code,
	}
}

func (e *HttpError) ToResponse() HttpResponse[any] {
	return NewResponseFromError(e)
}
