package models_test

import (
	"aws_challenge_pragma/domain/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser_GetTableName(t *testing.T) {
	user := models.User{}
	assert.Equal(t, "users", user.GetTableName())
}

func TestUser_GetPreloads(t *testing.T) {
	user := models.User{}
	assert.Equal(t, []string{}, user.GetPreloads())
}

func TestUser_GetPrimaryKey(t *testing.T) {
	user := models.User{}
	assert.Equal(t, "id", user.GetPrimaryKey())
}

func TestUser_Fields(t *testing.T) {
	id := uuid.New()
	user := models.User{
		Id:             id,
		DocumentNumber: 12345678,
		Name:           "John Doe",
		Email:          "john.doe@example.com",
	}

	assert.Equal(t, id, user.Id)
	assert.Equal(t, 12345678, user.DocumentNumber)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)
}
