package usecases_test

import (
	"errors"
	"testing"

	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserGetRepository struct {
	mock.Mock
}

func (m *MockUserGetRepository) Save(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserGetRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return []models.User{}, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}
func TestGetUsers_Success(t *testing.T) {
	mockRepo := new(MockUserGetRepository)
	useCase := usecases.NewGetUsers(mockRepo)

	users := []models.User{
		{Name: "John Doe", DocumentNumber: 30340703, Email: "john.doe@mail.com", Id: uuid.New()},
		{Name: "Jane Doe", DocumentNumber: 40450704, Email: "jane.doe@mail.com", Id: uuid.New()},
	}

	mockRepo.On("GetAll").Return(users, nil)

	result, err := useCase.Invoke()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, users, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_RepositoryError(t *testing.T) {
	mockRepo := new(MockUserGetRepository)
	useCase := usecases.NewGetUsers(mockRepo)

	repoErr := errors.New("repository error")
	mockRepo.On("GetAll").Return([]models.User{}, repoErr)

	result, err := useCase.Invoke()

	assert.Error(t, err)
	assert.Equal(t, repoErr, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_NilRepository(t *testing.T) {
	useCase := usecases.NewGetUsers(nil)

	result, err := useCase.Invoke()

	assert.Error(t, err)
	assert.Empty(t, result)
}
