package alc_result

type CommonErr Result // code=xxxxx, 业务错误，必须是5位或以上

func (e *CommonErr) Error() string {
	return e.Msg
}

func (e *CommonErr) Prefix(s string) *CommonErr {
	return &CommonErr{
		e.Code,
		s + e.Msg,
		e.Data,
	}
}

func (e *CommonErr) Suffix(s string) *CommonErr {
	return &CommonErr{
		e.Code,
		e.Msg + s,
		e.Data,
	}
}

func NewCommonErr(code int, msg string) *CommonErr {
	return &CommonErr{
		code,
		msg,
		nil,
	}
}
