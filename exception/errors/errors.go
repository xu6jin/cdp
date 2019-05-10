package errors

import (
	"errors"
	"fmt"
)

func New(f string, v ...interface{}) error {
	return errors.New(fmt.Sprintf(f, v...))
}
