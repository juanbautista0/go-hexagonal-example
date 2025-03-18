package user_repository_impl

import (
	"aws_challenge_pragma/domain/models"
	"aws_challenge_pragma/domain/repositories"
	"sync"

	"github.com/google/uuid"
)

var NewUserMemoryRepositoryImpl func() repositories.UserRepository = userMemoryRepositoryImpl

type UserMemoryRepositoryImpl struct {
	users map[string]models.User
	mu    sync.Mutex
}

func userMemoryRepositoryImpl() repositories.UserRepository {
	repo := &UserMemoryRepositoryImpl{
		users: make(map[string]models.User),
	}

	repo.users["46c819eb-ca23-429b-8be7-1556fef77103"] = models.User{Id: uuid.MustParse("46c819eb-ca23-429b-8be7-1556fef77103"), Name: "Juan", Email: "juan@example.com", DocumentNumber: 30340703}
	repo.users["0b5a21fc-107f-4052-bd76-6092d2074a16"] = models.User{Id: uuid.MustParse("0b5a21fc-107f-4052-bd76-6092d2074a16"), Name: "Maria", Email: "maria@example.com", DocumentNumber: 94269405}

	return repo
}

func (u *UserMemoryRepositoryImpl) GetAll() ([]models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]models.User, 0, len(u.users))
	for _, user := range u.users {
		users = append(users, user)
	}

	return users, nil
}

func (u *UserMemoryRepositoryImpl) Save(user *models.User) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user.Id = uuid.New()

	u.users[user.Id.String()] = *user
	return user, nil
}

func (u *UserMemoryRepositoryImpl) Clear(user *models.User) {

}
