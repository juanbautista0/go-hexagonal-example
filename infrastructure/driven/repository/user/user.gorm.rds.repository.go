package user_repository_impl

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/repositories"
	"aws_challenge_pragma/infrastructure/client"

	"github.com/google/uuid"
)

type UserGormRdsRepositoryImpl struct {
	client     client.GormRdsClient
	preloads   []string
	primaryKey string
	tableName  string
}

func NewUserGormRdsRepositoryImpl(dabatabaseClient client.GormRdsClient) repositories.UserRepository {
	return &UserGormRdsRepositoryImpl{
		client:     dabatabaseClient,
		preloads:   models.User{}.GetPreloads(),
		primaryKey: models.User{}.GetPrimaryKey(),
		tableName:  models.User{}.GetTableName(),
	}

}

func (u *UserGormRdsRepositoryImpl) GetAll() ([]models.User, error) {
	users := []models.User{}

	conn, err := u.client.GetInstance()
	if err != nil {
		return users, err
	}

	if err := conn.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserGormRdsRepositoryImpl) Save(user *models.User) (*models.User, error) {

	conn, err := u.client.GetInstance()
	if err != nil {
		return user, err
	}
	if uuid.Nil == user.Id {
		user.Id = uuid.New()
	}

	err = conn.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
