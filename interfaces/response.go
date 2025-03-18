package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type DtoResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data,omitempty"`
}

type GenericChan[T any] chan T

const ContentType = "application/json; charset=UTF-8"

func LambdaJsonResponse(statusCode int, message string, data any) events.APIGatewayProxyResponse {
	headers := map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
	}

	responseBody, err := json.Marshal(DtoResponse{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	})

	if err != nil {
		errorResponse := map[string]interface{}{
			"message":     http.StatusText(http.StatusInternalServerError),
			"status_code": http.StatusInternalServerError,
			"data":        nil,
		}
		errorBody, _ := json.Marshal(errorResponse)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       string(errorBody),
			Headers:    headers,
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(responseBody),
		Headers:    headers,
	}
}
