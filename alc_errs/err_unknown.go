package alc_errs

import "net/http"

// Unknown
type Unknown struct {
	HttpErr
}

func NewUnknown(message string) *Unknown {
	r := new(Unknown)
	r.Status = http.StatusInternalServerError
	r.Code = http.StatusInternalServerError
	r.Message = message
	return r
}

func (e *Unknown) Error() string {
	return e.Message
}
