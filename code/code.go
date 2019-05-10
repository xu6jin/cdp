package code

import (
	"net/http"
)

type Code interface {
	Status() int
	String() string
}

type code struct {
	status int
	text   string
}

func (c *code) Status() int {
	return c.status
}
func (c *code) String() string {
	return c.text
}

var (
	CodeOK                  = &code{status: 0, text: http.StatusText(http.StatusOK)}
	CodeCreated             = &code{status: 0, text: http.StatusText(http.StatusCreated)}
	CodeUpdated             = &code{status: 0, text: "Updated"}
	CodeNoContent           = &code{status: 0, text: http.StatusText(http.StatusNoContent)}
	CodeBadRequest          = &code{status: http.StatusBadRequest, text: http.StatusText(http.StatusBadRequest)}
	CodeUnauthorized        = &code{status: http.StatusUnauthorized, text: http.StatusText(http.StatusUnauthorized)}
	CodeForbidden           = &code{status: http.StatusForbidden, text: http.StatusText(http.StatusForbidden)}
	CodeNotFound            = &code{status: http.StatusNotFound, text: http.StatusText(http.StatusNotFound)}
	CodeMethodNotAllowed    = &code{status: http.StatusMethodNotAllowed, text: http.StatusText(http.StatusMethodNotAllowed)}
	CodeInternalServerError = &code{status: http.StatusInternalServerError, text: http.StatusText(http.StatusInternalServerError)}
	CodeBadGateway          = &code{status: http.StatusBadGateway, text: http.StatusText(http.StatusBadGateway)}
	CodeGatewayTimeout      = &code{status: http.StatusGatewayTimeout, text: http.StatusText(http.StatusGatewayTimeout)}
)
