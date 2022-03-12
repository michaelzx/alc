package alc_result

type CommonErr Result // code=xxxxx, 业务错误，必须是5位或以上

func (e *CommonErr) Error() string {
	return e.Msg
}

func NewCommonErr(code int, msg string) *CommonErr {
	return &CommonErr{
		code,
		msg,
		nil,
	}
}
