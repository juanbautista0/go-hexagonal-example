package driving

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/usecases"
	err "aws_challenge_pragma/infrastructure/error"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func NewGetUsersLambdaAdapter(req events.APIGatewayProxyRequest, useCase *usecases.GetUsers) ([]models.User, *err.CustomError) {
	var invokeErr error
	var result []models.User

	if result, invokeErr = useCase.Invoke(); invokeErr != nil {
		return []models.User{}, err.NewCustomError(http.StatusInternalServerError, invokeErr.Error())
	}

	return result, nil
}
