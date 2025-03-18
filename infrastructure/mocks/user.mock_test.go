package mocks_test

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/infrastructure/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserMock(t *testing.T) {
	assert.NotNil(t, mocks.UserMock, "UserMock should not be nil")
	assert.IsType(t, models.User{}, mocks.UserMock, "UserMock should be of type models.User")
	assert.NotEqual(t, uuid.Nil, mocks.UserMock.Id, "UserMock Id should not be nil")
	assert.NotEmpty(t, mocks.UserMock.Name, "UserMock Name should not be empty")
	assert.NotZero(t, mocks.UserMock.DocumentNumber, "UserMock DocumentNumber should not be zero")
	assert.NotEmpty(t, mocks.UserMock.Email, "UserMock Email should not be empty")
}

func TestUsersMock(t *testing.T) {
	assert.NotNil(t, mocks.UsersMock, "UsersMock should not be nil")
	assert.Greater(t, len(mocks.UsersMock), 0, "UsersMock should contain at least one user")
	for _, user := range mocks.UsersMock {
		assert.IsType(t, models.User{}, user, "Each entry in UsersMock should be of type models.User")
		assert.NotEqual(t, uuid.Nil, user.Id, "User Id should not be nil")
		assert.NotEmpty(t, user.Name, "User Name should not be empty")
		assert.NotZero(t, user.DocumentNumber, "User DocumentNumber should not be zero")
		assert.NotEmpty(t, user.Email, "User Email should not be empty")
	}
}
