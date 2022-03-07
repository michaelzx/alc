package alc_result

import "net/http"

type UnauthorizedErr CommonErr // code=401, 3位留给http status
type ForbiddenErr CommonErr    // code=403, 3位留给http status
type NotFoundErr CommonErr     // code=404, 3位留给http status

func NewUnauthorizedErr() *UnauthorizedErr {
	return &UnauthorizedErr{
		http.StatusUnauthorized,
		"Unauthorized",
		nil,
	}
}

func NewForbiddenErr() *ForbiddenErr {
	return &ForbiddenErr{
		http.StatusForbidden,
		"Forbidden",
		nil,
	}
}

func NewNotFoundErr() *NotFoundErr {
	return &NotFoundErr{
		http.StatusNotFound,
		"Not Found",
		nil,
	}
}
