package exception

import "github.com/xu6jin/cdp/exception/errors"

type Exception interface {
	Status() int
	Errors() []error
	Format() []string
}

type exception struct {
	status int
	errs   []error
}

func (e *exception) Errors() []error {
	return e.errs
}

func (e *exception) Status() int {
	return e.status
}

func (e *exception) Format() []string {
	format := make([]string, 0)
	for _, err := range e.errs {
		format = append(format, err.Error())
	}
	return format
}

func New(status int, errs ...error) Exception {
	es := make([]error, 0)
	for _, err := range errs {
		if _, ok := err.(errors.ErrorObject); ok {
			es = append(es, err)
			continue
		}
		es = append(es, errors.New(errors.ErrorCodeUnknown, err))
	}
	return &exception{status: status, errs: errs}
}
