package exception

import (
	"github.com/xu6jin/cdp/code"
)

type Exception interface {
	error
	Code() code.Code
	Errors() []error
}

type exception struct {
	code code.Code
	errs []error
}

func (e *exception) Error() string {
	return e.code.Text()
}

func (e *exception) Code() code.Code {
	return e.code
}

func (e *exception) Errors() []error {
	return e.errs
}

func New(c code.Code, errs ...error) Exception {
	return &exception{code: c, errs: errs}
}
