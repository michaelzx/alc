package alc_errs

import (
	"fmt"
	"strings"
)

type IBizErr interface {
}

type BizErr struct {
	Code    int
	Message string
}
type HttpErr struct {
	Status int
	BizErr
}

func Wrap(e error, msgArgs ...string) error {
	msg := strings.Join(msgArgs, " ")
	return fmt.Errorf("%s: %w", msg, e)
}
