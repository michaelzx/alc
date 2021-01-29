package alc_errs

import "net/http"

// Unauthorized
type Unauthorized struct {
	HttpErr
}

func NewUnauthorized(message string) *Unauthorized {
	r := new(Unauthorized)
	r.Status = http.StatusUnauthorized
	r.Code = http.StatusUnauthorized
	r.Message = message
	return r
}

func (e *Unauthorized) Error() string {
	return e.Message
}
