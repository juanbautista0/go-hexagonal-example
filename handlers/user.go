package handlers

import (
	"aws_challenge_pragma/domain/repositories"
	"aws_challenge_pragma/domain/usecases"
	"aws_challenge_pragma/infrastructure/client"
	"aws_challenge_pragma/infrastructure/driving"
	"aws_challenge_pragma/infrastructure/logger"
	"aws_challenge_pragma/interfaces"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type UserHandlerDependencies struct {
	UserRepository func(dabatabaseClient client.GormRdsClient) repositories.UserRepository
	RdsClient      func() (client.GormRdsClient, error)
}

var appLogger logger.AppLogger = logger.AppLogger{}

func UserHandler(req events.APIGatewayProxyRequest, repository repositories.UserRepository) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return CreateUser(req, repository)
	case "GET":
		return GetUsers(req, repository)
	default:
		return interfaces.LambdaJsonResponse(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil), errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}
}

func CreateUser(req events.APIGatewayProxyRequest, repository repositories.UserRepository) (events.APIGatewayProxyResponse, error) {
	useCase := usecases.NewCreateUser(repository)

	if _, err := driving.NewCreateUserLambdaAdapter(req, useCase); err != nil {
		appLogger.Error().
			SetCode("Handler.CreateUser").
			SetDetail(err.Detail).
			SetMessage(err.Message).
			SetMetadata(map[string]interface{}{
				"error": err.Error(),
			}).Write()
		return interfaces.LambdaJsonResponse(err.Code, err.Message, err), nil
	}

	return interfaces.LambdaJsonResponse(http.StatusCreated, http.StatusText(http.StatusCreated), nil), nil
}

func GetUsers(req events.APIGatewayProxyRequest, repository repositories.UserRepository) (events.APIGatewayProxyResponse, error) {
	useCase := usecases.NewGetUsers(repository)

	users, err := driving.NewGetUsersLambdaAdapter(req, useCase)
	if err != nil {
		appLogger.Error().
			SetCode("Handler.CreateUser").
			SetDetail(err.Detail).
			SetMessage(err.Message).
			SetMetadata(map[string]interface{}{
				"error": err.Error(),
			}).Write()
		return interfaces.LambdaJsonResponse(err.Code, err.Message, err), errors.New(err.Error())
	}

	return interfaces.LambdaJsonResponse(http.StatusOK, http.StatusText(http.StatusOK), users), nil
}

func LambdaHandler(dependencies *UserHandlerDependencies) func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		appLogger.Info().
			SetCode("LambdaHandler.Users").
			SetDetail("calling handler method").
			SetMessage("eventsAPIGatewayProxyRequest").
			SetMetadata(map[string]interface{}{
				"eventsAPIGatewayProxyRequest": req,
			}).Write()

		rdsClient, err := dependencies.RdsClient()
		if err != nil {
			appLogger.Error().
				SetCode("LambdaHandler").
				SetDetail("LambdaHandler:rdsClient, err := dependencies.RdsClient()").
				SetMessage("instance of rdsClient").
				SetMetadata(map[string]interface{}{
					"error": err.Error(),
				}).Write()
			return interfaces.LambdaJsonResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil), errors.New(http.StatusText(http.StatusInternalServerError))
		}

		useRepository := dependencies.UserRepository(rdsClient)
		path := strings.ToLower(req.Path)

		if req.HTTPMethod == "OPTIONS" {
			return interfaces.LambdaJsonResponse(http.StatusOK, http.StatusText(http.StatusOK), nil), nil
		}

		switch {
		case strings.HasPrefix(path, "/v1/users"):
			return UserHandler(req, useRepository)
		default:
			return interfaces.LambdaJsonResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil), errors.New(http.StatusText(http.StatusNotFound))
		}
	}
}
