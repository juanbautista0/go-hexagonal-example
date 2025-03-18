package user_repository_impl_test

import (
	"errors"
	"fmt"
	"testing"

	"aws_challenge_pragma/domain/models"
	user_repository_impl "aws_challenge_pragma/infrastructure/driven/repository/user"
	"aws_challenge_pragma/infrastructure/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockGormRdsClient = mocks.MockGormRdsClient

func TestGetAll_Success(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)

	mockUsers := mocks.UsersMock

	rows := sqlmock.NewRows([]string{"id", "name", "document_number", "email"}).
		AddRow(mockUsers[0].Id, mockUsers[0].Name, mockUsers[0].DocumentNumber, mockUsers[0].Email).
		AddRow(mockUsers[1].Id, mockUsers[1].Name, mockUsers[1].DocumentNumber, mockUsers[1].Email)

	mock.ExpectQuery(`SELECT \* FROM "users"`).WillReturnRows(rows)

	mockClient := new(MockGormRdsClient)
	mockClient.On("GetInstance").Return(db, nil)

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockClient)
	users, err := repo.GetAll()
	fmt.Println(len(users))
	assert.NoError(t, err)
	assert.Equal(t, len(mockUsers), len(users))
	mockClient.AssertExpectations(t)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_ErrorFetchingInstance(t *testing.T) {
	mockClient := new(MockGormRdsClient)
	mockClient.On("GetInstance").Return((*gorm.DB)(nil), errors.New("failed to connect to DB"))

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockClient)

	users, err := repo.GetAll()

	assert.Error(t, err)
	assert.Empty(t, users)
	mockClient.AssertExpectations(t)
}

func TestSave_Success(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)

	mockUser := models.User{
		Id:             uuid.New(),
		Name:           "Jane Doe",
		DocumentNumber: 12345678,
		Email:          "jane.doe@mail.com",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mockClient := new(MockGormRdsClient)
	mockClient.On("GetInstance").Return(db, nil)

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockClient)
	savedUser, err := repo.Save(&mockUser)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.Id, savedUser.Id)
	assert.Equal(t, mockUser.Name, savedUser.Name)
	mockClient.AssertExpectations(t)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSave_ErrorFetchingInstance(t *testing.T) {
	mockClient := new(MockGormRdsClient)
	mockClient.On("GetInstance").Return((*gorm.DB)(nil), errors.New("failed to connect to DB"))

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockClient)
	mockUser := models.User{Id: uuid.New(), Name: "Jane Doe", DocumentNumber: 12345678, Email: "jane.doe@mail.com"}

	savedUser, err := repo.Save(&mockUser)

	assert.Error(t, err)
	assert.Equal(t, &mockUser, savedUser)
	mockClient.AssertExpectations(t)
}

func TestSave_ErrorInsertingUser(t *testing.T) {
	db, mock, err := mocks.GetDatabaseMock()
	assert.NoError(t, err)

	mockUser := models.User{Id: uuid.New(), Name: "Jane Doe", DocumentNumber: 12345678, Email: "jane.doe@mail.com"}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).WillReturnError(errors.New("insert failed"))
	mock.ExpectRollback()

	mockClient := new(MockGormRdsClient)
	mockClient.On("GetInstance").Return(db, nil)

	repo := user_repository_impl.NewUserGormRdsRepositoryImpl(mockClient)
	savedUser, err := repo.Save(&mockUser)

	assert.Error(t, err)
	assert.Equal(t, &mockUser, savedUser)
	mockClient.AssertExpectations(t)
	assert.NoError(t, mock.ExpectationsWereMet())
}
