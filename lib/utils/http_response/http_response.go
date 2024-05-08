package http_response

type HttpResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewHttpResponseWithError[T *any](message string) HttpResponse[T] {
	return HttpResponse[T]{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}
