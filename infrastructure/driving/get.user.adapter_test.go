package driving_test

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/usecases"
	"aws_challenge_pragma/infrastructure/driving"
	"aws_challenge_pragma/infrastructure/mocks"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestNewGetUsersHttpAdapter_Success(t *testing.T) {
	users := mocks.UsersMock
	mockRepo := new(MockUserRepository)
	mockRepo.On("GetAll").Return(users, nil)

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/v1/users",
	}

	mockUseCase := usecases.NewGetUsers(mockRepo)
	response, err := driving.NewGetUsersLambdaAdapter(req, mockUseCase)

	assert.Nil(t, err)
	assert.Equal(t, len(users), len(response))
}

func TestNewGetUsersHttpAdapter_Fail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("GetAll").Return([]models.User{}, errors.New("Fail to get"))

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/v1/users",
	}

	mockUseCase := usecases.NewGetUsers(mockRepo)
	response, err := driving.NewGetUsersLambdaAdapter(req, mockUseCase)

	assert.NotNil(t, err)
	assert.Equal(t, 0, len(response))
}
