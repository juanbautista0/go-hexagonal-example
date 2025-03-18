package usecases

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/repositories"
	"errors"
)

type CreateUser struct {
	UserRepository repositories.UserRepository
}

func NewCreateUser(userRepository repositories.UserRepository) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
	}
}

func (instance *CreateUser) Invoke(user *models.User) (*models.User, error) {
	if user == nil {
		return user, errors.New("user cannot be nil")
	}

	if instance.UserRepository == nil {
		return nil, errors.New("repository cannot be nil")
	}

	return instance.UserRepository.Save(user)
}
