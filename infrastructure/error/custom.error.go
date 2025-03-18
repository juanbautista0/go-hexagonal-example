package err

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Code    int
	Message string
	Detail  string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("CustomError code: %d: message: %s - detail: %s", e.Code, e.Message, e.Detail)
}

func NewCustomError(code int, detail string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: http.StatusText(code),
		Detail:  detail,
	}
}
