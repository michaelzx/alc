package alc_result

type Result struct {
	Code int         `json:"code"` //
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (e *Result) WithData(data interface{}) *Result {
	e.Data = data
	return e
}
