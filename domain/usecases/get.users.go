package usecases

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/repositories"
	"errors"
)

type GetUsers struct {
	UserRepository repositories.UserRepository
}

func NewGetUsers(userRepository repositories.UserRepository) *GetUsers {
	return &GetUsers{
		UserRepository: userRepository,
	}
}

func (instance *GetUsers) Invoke() ([]models.User, error) {
	if instance.UserRepository == nil {
		return []models.User{}, errors.New("repository cannot be nil")
	}

	return instance.UserRepository.GetAll()
}
