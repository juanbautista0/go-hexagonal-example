package usecases_test

import (
	"errors"
	"testing"

	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockUserCreateRepository struct {
	mock.Mock
}

func (m *MockUserCreateRepository) Save(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserCreateRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestCreateUser_Invoke_Success(t *testing.T) {
	mockRepo := new(MockUserCreateRepository)
	useCase := usecases.NewCreateUser(mockRepo)

	user := &models.User{Name: "John Doe", DocumentNumber: 30340703, Email: "john.doe@mail.com", Id: uuid.New()}
	mockRepo.On("Save", user).Return(user, nil)

	createdUser, err := useCase.Invoke(user)

	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Invoke_Error(t *testing.T) {
	mockRepo := new(MockUserCreateRepository)
	useCase := usecases.NewCreateUser(mockRepo)

	user := &models.User{Name: "John Doe", DocumentNumber: 30340703, Email: "john.doe@mail.com", Id: uuid.New()}
	expectedErr := errors.New("failed to save user")
	mockRepo.On("Save", user).Return((*models.User)(nil), expectedErr)

	createdUser, err := useCase.Invoke(user)

	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Invoke_NilUser(t *testing.T) {
	mockRepo := new(MockUserCreateRepository)
	useCase := usecases.NewCreateUser(mockRepo)

	_, err := useCase.Invoke(nil)

	assert.Error(t, err, "user cannot be nil")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCreateUser_Constructor(t *testing.T) {
	mockRepo := new(MockUserCreateRepository)
	createUserInstance := usecases.NewCreateUser(mockRepo)

	assert.NotNil(t, createUserInstance)
	assert.Equal(t, mockRepo, createUserInstance.UserRepository)
}

func TestCreateUser_NilRepository(t *testing.T) {
	useCase := usecases.NewCreateUser(nil)
	user := &models.User{}

	result, err := useCase.Invoke(user)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "repository cannot be nil", err.Error())
}
