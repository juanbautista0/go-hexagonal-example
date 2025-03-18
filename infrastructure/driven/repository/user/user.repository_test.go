package user_repository_impl_test

import (
	"aws_challenge_pragma/domain/models"
	user_repository_impl "aws_challenge_pragma/infrastructure/driven/repository/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserMemoryRepositoryImpl_GetAll(t *testing.T) {
	repo := user_repository_impl.NewUserMemoryRepositoryImpl()

	users, err := repo.GetAll()
	assert.NoError(t, err, "GetAll should not return an error")
	assert.Len(t, users, 2, "GetAll should return exactly 2 users")
}

func TestUserMemoryRepositoryImpl_Save(t *testing.T) {
	repo := user_repository_impl.NewUserMemoryRepositoryImpl()

	newUser := &models.User{Name: "Carlos", Email: "carlos@example.com", DocumentNumber: 12345678, Id: uuid.New()}
	savedUser, err := repo.Save(newUser)
	assert.NoError(t, err, "Save should not return an error")
	assert.NotNil(t, savedUser, "Saved user should not be nil")
	assert.NotEqual(t, uuid.Nil, savedUser.Id, "Saved user should have a valid UUID")

	users, _ := repo.GetAll()
	assert.Len(t, users, 3, "GetAll should return 3 users after adding a new one")
}
