package alc_errs

import "net/http"

// NotFound 用户未授权
type NotFound struct {
	HttpErr
}

const notFoundMessage = "Page Not Found"

func NewNotFound() *NotFound {
	r := new(NotFound)
	r.Status = http.StatusNotFound
	r.Code = http.StatusNotFound
	r.Message = notFoundMessage
	return r
}

func (e *NotFound) Error() string {
	return e.Message
}
