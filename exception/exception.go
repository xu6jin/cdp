package exception

import (
	"github.com/xu6jin/cdp/code"
)

type Exception interface {
	error
	Code() code.Code
}

type exception struct {
	code code.Code
	err  error
}

func (e *exception) Error() string {
	return e.err.Error()
}

func (e *exception) Code() code.Code {
	return e.code
}

func New(c code.Code, err error) Exception {
	return &exception{code: c, err: err}
}
