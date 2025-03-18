package mocks

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return db, mock, nil
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user interface{}) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUsers() ([]interface{}, error) {
	args := m.Called()
	return args.Get(0).([]interface{}), args.Error(1)
}
