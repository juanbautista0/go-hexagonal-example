package mocks_test

import (
	"aws_challenge_pragma/infrastructure/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabaseMock(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err, "GetDatabaseMock should not return an error")
	assert.NotNil(t, db, "Database instance should not be nil")
	assert.NotNil(t, mock, "SQL mock instance should not be nil")
}

func TestMockUserRepository_CreateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	user := struct{ Name string }{Name: "John Doe"}

	mockRepo.On("CreateUser", user).Return(nil)

	err := mockRepo.CreateUser(user)
	assert.NoError(t, err, "CreateUser should not return an error")
	mockRepo.AssertExpectations(t)
}

func TestMockUserRepository_GetUsers(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	expectedUsers := []interface{}{struct{ Name string }{Name: "John Doe"}}
	mockRepo.On("GetUsers").Return(expectedUsers, nil)

	users, err := mockRepo.GetUsers()
	assert.NoError(t, err, "GetUsers should not return an error")
	assert.Equal(t, expectedUsers, users, "Returned users should match expected users")
	mockRepo.AssertExpectations(t)
}
