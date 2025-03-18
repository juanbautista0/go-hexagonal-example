package interfaces_test

import (
	"aws_challenge_pragma/interfaces"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaJsonResponse(t *testing.T) {
	data := map[string]string{"key": "value"}
	resp := interfaces.LambdaJsonResponse(http.StatusOK, "Success", data)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Headers["Content-Type"])

	var body interfaces.DtoResponse
	err := json.Unmarshal([]byte(resp.Body), &body)
	assert.NoError(t, err)
	assert.Equal(t, "Success", body.Message)
	assert.Equal(t, http.StatusOK, body.StatusCode)

	dataMap, ok := body.Data.(map[string]interface{})
	assert.True(t, ok, "Expected body.Data to be a map[string]interface{}")

	expectedData := map[string]interface{}{"key": "value"}
	assert.Equal(t, expectedData, dataMap)
}

func TestLambdaJsonResponse_Error(t *testing.T) {
	resp := interfaces.LambdaJsonResponse(http.StatusOK, "Success", make(chan int)) // Invalid data

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var body interfaces.DtoResponse
	err := json.Unmarshal([]byte(resp.Body), &body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), body.Message)
	assert.Equal(t, http.StatusInternalServerError, body.StatusCode)
	assert.Nil(t, body.Data)
}
