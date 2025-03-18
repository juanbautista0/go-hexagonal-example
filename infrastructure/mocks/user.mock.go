package mocks

import (
	"aws_challenge_pragma/domain/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var UserMock = models.User{Id: uuid.New(), Name: "John Doe", DocumentNumber: 30340703, Email: "john.doe@mail.com"}
var UsersMock = []models.User{UserMock, {Id: uuid.New(), Name: "Johana Doe", DocumentNumber: 9452168, Email: "johana.doe@mail.com"}}

type MockGormRdsClient struct {
	mock.Mock
}

func (m *MockGormRdsClient) GetInstance() (*gorm.DB, error) {
	args := m.Called()
	return args.Get(0).(*gorm.DB), args.Error(1)
}
