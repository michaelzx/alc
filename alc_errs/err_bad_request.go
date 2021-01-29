package alc_errs

import "net/http"

const CommonBizCode = 10000

// BadRequest
type BadRequest struct {
	HttpErr
}

func NewBadRequest(code int, message string) *BadRequest {
	r := new(BadRequest)
	r.Status = http.StatusBadRequest
	r.Code = code
	r.Message = message
	return r
}
func CommonError(message string) *BadRequest {
	r := new(BadRequest)
	r.Status = http.StatusBadRequest
	r.Code = CommonBizCode
	r.Message = message
	return r
}

func (e *BadRequest) Error() string {
	return e.Message
}

func (e *BadRequest) Suffix(msg string) *BadRequest {
	r := new(BadRequest)
	r.Status = http.StatusBadRequest
	r.Code = e.Code
	r.Message = e.Message + msg
	return r
}

func (e *BadRequest) Prefix(msg string) *BadRequest {
	r := new(BadRequest)
	r.Status = http.StatusBadRequest
	r.Code = e.Code
	r.Message = msg + e.Message
	return r
}
