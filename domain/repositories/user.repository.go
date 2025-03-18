package repositories

import "aws_challenge_pragma/domain/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	Save(user *models.User) (*models.User, error)
}
