package alc_result

type CommonErr Result // code=xxxxx, 业务错误，必须是5位或以上

func (e *CommonErr) Error() string {
	return e.Msg
}

func (e *CommonErr) Prefix(s string) *CommonErr {
	e.Msg = s + e.Msg
	return e
}

func (e *CommonErr) Suffix(s string) *CommonErr {
	e.Msg = e.Msg + s
	return e
}

func NewCommonErr(code int, msg string) *CommonErr {
	return &CommonErr{
		code,
		msg,
		nil,
	}
}
