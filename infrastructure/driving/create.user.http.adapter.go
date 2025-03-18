package driving

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/usecases"
	err "aws_challenge_pragma/infrastructure/error"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func NewCreateUserLambdaAdapter(req events.APIGatewayProxyRequest, useCase *usecases.CreateUser) (bool, *err.CustomError) {
	if req.Body == "" {
		return false, err.NewCustomError(http.StatusBadRequest, "invalid or empty request body")
	}

	var input models.User
	if decodeErr := json.Unmarshal([]byte(req.Body), &input); decodeErr != nil {
		return false, err.NewCustomError(http.StatusBadRequest, "invalid JSON format")
	}

	_, invokeErr := useCase.Invoke(&input)
	if invokeErr != nil {
		return false, err.NewCustomError(http.StatusInternalServerError, invokeErr.Error())
	}

	return true, nil
}
