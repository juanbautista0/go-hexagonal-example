package main

import (
	"aws_challenge_pragma/handlers"
	"aws_challenge_pragma/infrastructure/client"
	"aws_challenge_pragma/infrastructure/config"
	user_repository_impl "aws_challenge_pragma/infrastructure/driven/repository/user"

	"github.com/aws/aws-lambda-go/lambda"
)

func InitDependencies() *handlers.UserHandlerDependencies {
	config.LoadConfig()
	return &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      client.NewRdsMySQLGormClient,
	}
}

func main() {
	dependencies := InitDependencies()
	lambda.Start(handlers.LambdaHandler(dependencies))
}
