package http_response

type HttpResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type CursorPage[T any] struct {
	Data   []T    `json:"data"`
	Cursor string `json:"cursor"`
}

type HttpPagedResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	CursorPage[T]
}

func NewHttpResponseWithError[T *any](message string) HttpResponse[T] {
	return HttpResponse[T]{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}

func NewResponseFromError(err *HttpError) HttpResponse[any] {
	return HttpResponse[any]{
		Status:  false,
		Message: err.Message,
		Data:    nil,
	}
}

func NewSuccessResponse[T any](data T) HttpResponse[T] {
	return HttpResponse[T]{
		Status:  true,
		Message: "success",
		Data:    data,
	}
}

func NewPagedResponse[T any](page CursorPage[T]) HttpPagedResponse[T] {
	return HttpPagedResponse[T]{
		Status:     true,
		Message:    "success",
		CursorPage: page,
	}
}
