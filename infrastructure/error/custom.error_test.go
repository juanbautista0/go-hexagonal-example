package err_test

import (
	"net/http"
	"testing"

	err "aws_challenge_pragma/infrastructure/error"

	"github.com/stretchr/testify/assert"
)

func TestCustomError_Error(t *testing.T) {
	err := err.CustomError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Detail:  "Invalid input data",
	}

	expected := "CustomError code: 400: message: Bad Request - detail: Invalid input data"
	assert.Equal(t, expected, err.Error())
}

func TestNewCustomError(t *testing.T) {
	err := err.NewCustomError(http.StatusNotFound, "User not found")

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Code)
	assert.Equal(t, "Not Found", err.Message)
	assert.Equal(t, "User not found", err.Detail)
}
