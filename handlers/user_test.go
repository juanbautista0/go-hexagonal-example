package handlers_test

import (
	"aws_challenge_pragma/handlers"
	"aws_challenge_pragma/infrastructure/client"
	user_repository_impl "aws_challenge_pragma/infrastructure/driven/repository/user"
	"aws_challenge_pragma/infrastructure/mocks"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockGormRdsClient = mocks.MockGormRdsClient
type MockUserRepository = mocks.MockUserRepository

func getMockRdsClient(db *gorm.DB) func() (client.GormRdsClient, error) {
	return func() (client.GormRdsClient, error) {
		mockClient := new(MockGormRdsClient)
		mockClient.On("GetInstance").Return(db, nil)
		return mockClient, nil
	}
}

func TestUserHandler_Post_Success(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)

	mockRdsClient := getMockRdsClient(db)

	user := mocks.UserMock
	body, err := json.Marshal(user)
	assert.NoError(t, err)
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users" \(.*\) VALUES \(.*\)`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	request := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       string(body),
		Path:       "/v1/users",
	}

	dependencies := &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      mockRdsClient,
	}

	response, err := handlers.LambdaHandler(dependencies)(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, 201, response.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestUserHandler_Post_Fail(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient := getMockRdsClient(db)

	user := mocks.UserMock
	body, _ := json.Marshal(user)

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users" \(.*\) VALUES \(.*\)`).
		WillReturnError(errors.New("database error"))
	mock.ExpectRollback()

	request := events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: string(body), Path: "/v1/users"}
	dependencies := &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      mockRdsClient,
	}

	response, err := handlers.LambdaHandler(dependencies)(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserHandler_Get_Success(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient := getMockRdsClient(db)

	rows := sqlmock.NewRows([]string{"id", "name", "document_number", "email"})
	for _, user := range mocks.UsersMock {
		rows.AddRow(user.Id, user.Name, user.DocumentNumber, user.Email)
	}

	mock.ExpectQuery(`SELECT\s+\*?\s+FROM\s+"users"`).WillReturnRows(rows)

	request := events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/v1/users"}
	dependencies := &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      mockRdsClient,
	}

	response, err := handlers.LambdaHandler(dependencies)(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserHandler_Get_Empty(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient := getMockRdsClient(db)

	mock.ExpectQuery(`SELECT\s+\*?\s+FROM\s+"users"`).WillReturnRows(sqlmock.NewRows([]string{}))

	request := events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/v1/users"}
	dependencies := &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      mockRdsClient,
	}

	response, err := handlers.LambdaHandler(dependencies)(context.Background(), request)

	expectedBody := `{"message":"OK","status_code":200,"data":[]}`
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.JSONEq(t, expectedBody, response.Body)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUsersHandler_Adapter_Error(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient, _ := getMockRdsClient(db)()

	mock.ExpectQuery(`SELECT\s+\*?\s+FROM\s+"users"`).WillReturnError(errors.New("database error"))

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockRdsClient)
	request := events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/v1/users"}
	result, err := handlers.GetUsers(request, repo)
	assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	assert.NotNil(t, err)
}

func TestUserHandler_Method_No_Allowed(t *testing.T) {
	db, _, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient, _ := getMockRdsClient(db)()

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockRdsClient)
	request := events.APIGatewayProxyRequest{HTTPMethod: "X_DELETE", Path: "/v1/users"}
	result, err := handlers.UserHandler(request, repo)
	assert.Equal(t, http.StatusMethodNotAllowed, result.StatusCode)
	assert.NotNil(t, err)
}

func TestLambdaHandler_Invalid_Path(t *testing.T) {
	db, _, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient := getMockRdsClient(db)

	dependencies := &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      mockRdsClient,
	}
	request := events.APIGatewayProxyRequest{HTTPMethod: "X_DELETE", Path: "/v1.2/users"}
	result, err := handlers.LambdaHandler(dependencies)(context.Background(), request)
	assert.Equal(t, http.StatusNotFound, result.StatusCode)
	assert.NotNil(t, err)
}

func TestLambdaHandler_Http_Method_OPTIONS(t *testing.T) {
	db, _, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)
	mockRdsClient := getMockRdsClient(db)

	dependencies := &handlers.UserHandlerDependencies{
		UserRepository: user_repository_impl.NewUserGormRdsRepositoryImpl,
		RdsClient:      mockRdsClient,
	}
	request := events.APIGatewayProxyRequest{HTTPMethod: "OPTIONS", Path: "/v1/users"}
	result, err := handlers.LambdaHandler(dependencies)(context.Background(), request)
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Nil(t, err)
}
