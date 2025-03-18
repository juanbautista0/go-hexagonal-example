package driving_test

import (
	"encoding/json"
	"errors"
	"testing"

	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/usecases"
	"aws_challenge_pragma/infrastructure/driving"
	"aws_challenge_pragma/infrastructure/mocks"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateUserRepository struct {
	mock.Mock
}

func (m *MockCreateUserRepository) Save(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCreateUserRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestNewCreateUserHttpApdater_Success(t *testing.T) {
	user := mocks.UserMock
	body, err := json.Marshal(user)
	assert.NoError(t, err)

	mockRepo := new(MockCreateUserRepository)
	mockRepo.On("Save", mock.AnythingOfType("*models.User")).Return(&user, nil)

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       string(body),
		Path:       "/v1/users",
	}

	mockUseCase := usecases.NewCreateUser(mockRepo)
	success, err := driving.NewCreateUserLambdaAdapter(req, mockUseCase)
	assert.Equal(t, true, success)
	assert.Nil(t, err)
}

func TestNewCreateUserHttpApdater_Fail(t *testing.T) {
	user := mocks.UserMock
	body, err := json.Marshal(user)
	assert.NoError(t, err)

	mockRepo := new(MockCreateUserRepository)
	mockRepo.On("Save", mock.AnythingOfType("*models.User")).Return(&user, errors.New("Fail to save"))

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       string(body),
		Path:       "/v1/users",
	}

	mockUseCase := usecases.NewCreateUser(mockRepo)
	success, err := driving.NewCreateUserLambdaAdapter(req, mockUseCase)
	assert.Equal(t, false, success)
	assert.NotNil(t, err)
}

func TestNewCreateUserHttpApdater_Fail_Invalid_Body(t *testing.T) {
	user := mocks.UserMock

	mockRepo := new(MockCreateUserRepository)
	mockRepo.On("Save", mock.AnythingOfType("*models.User")).Return(&user, errors.New("Fail to save"))

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       "<{}>",
		Path:       "/v1/users",
	}

	mockUseCase := usecases.NewCreateUser(mockRepo)
	success, err := driving.NewCreateUserLambdaAdapter(req, mockUseCase)
	assert.Equal(t, false, success)
	assert.NotNil(t, err)
}

func TestNewCreateUserHttpApdater_Fail_Empty_Body(t *testing.T) {
	user := mocks.UserMock

	mockRepo := new(MockCreateUserRepository)
	mockRepo.On("Save", mock.AnythingOfType("*models.User")).Return(&user, errors.New("Fail to save"))

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       "",
		Path:       "/v1/users",
	}

	mockUseCase := usecases.NewCreateUser(mockRepo)
	success, err := driving.NewCreateUserLambdaAdapter(req, mockUseCase)
	assert.Equal(t, false, success)
	assert.NotNil(t, err)
}
