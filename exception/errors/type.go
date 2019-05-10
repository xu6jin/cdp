package errors

import "fmt"

type ErrorType uint8

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeField
)

func New(errorType ErrorType, format string, a ...interface{}) ErrorObject {
	if ErrorTypeUnknown > errorType || errorType > ErrorTypeField {
		errorType = ErrorTypeUnknown
	}
	return newErrorObject(errorType, fmt.Sprintf(format, a...))
}

func NewUnknownError(format string, a ...interface{}) ErrorObject {
	return newErrorObject(ErrorTypeUnknown, fmt.Sprintf(format, a...))
}

func NewFieldError(field string, value interface{}, msgs ...interface{}) ErrorObject {
	err := newErrorObject(ErrorTypeField, fmt.Sprintf("field error: %s", field))
	err.Set("field", field).Set("value", value)
	if len(msgs) > 0 {
		err.Set("text", fmt.Sprintf(fmt.Sprintf("%v", msgs[0]), msgs[1:]...))
	}
	return err
}
