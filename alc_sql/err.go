package alc_sql

import "fmt"

type Err struct {
	msg string
}

func NewErr(msg string) *Err {
	return &Err{msg: msg}
}

func (e *Err) Error() string {
	return fmt.Sprintf("alc_sql error: %s", e.msg)
}
