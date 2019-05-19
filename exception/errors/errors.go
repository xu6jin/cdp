package errors

import (
	"fmt"
)

type ErrorObject interface {
	error
	Code() ErrorCode
	Set(string, interface{}) ErrorObject
	Get(string) interface{}
}

type errorObject struct {
	err      error
	code     ErrorCode
	metadata map[string]interface{}
}

func (e *errorObject) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.err.Error())
}

func (e *errorObject) Code() ErrorCode {
	return e.code
}

func (e *errorObject) Set(k string, v interface{}) ErrorObject {
	if e.metadata == nil {
		e.metadata = make(map[string]interface{}, 0)
	}
	if v != nil {
		e.metadata[k] = v
	}
	return e
}

func (e *errorObject) Get(k string) interface{} {
	if v, ok := e.metadata[k]; ok {
		return v
	}
	return nil
}

func New(code ErrorCode, v ...interface{}) ErrorObject {
	format := "unknown error"
	a := make([]interface{}, 0)
	if len(v) > 0 {
		format = fmt.Sprint(v[0])
	}
	if len(v) > 1 {
		a = v[1:]
	}
	return &errorObject{err: fmt.Errorf(format, a...), code: code}
}
