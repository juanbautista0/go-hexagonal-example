package client_test

import (
	"aws_challenge_pragma/infrastructure/client"
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de la función de autenticación
type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) BuildAuthToken(ctx context.Context, endpoint, region, username string, creds aws.CredentialsProvider, optFns ...func(options *auth.BuildAuthTokenOptions)) (string, error) {
	args := m.Called(ctx, endpoint, region, username, creds)
	return args.String(0), args.Error(1)
}

// Mock para simular error en configuración AWS
type MockConfigLoader struct{}

func (m *MockConfigLoader) LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	return aws.Config{}, errors.New("error loading AWS config")
}

func resetSingleton() {
	client.ResetRdsGormClient()
}

func TestNewRdsMySQLGormClient_WithMockAuth(t *testing.T) {
	resetSingleton()

	clientInstance, err := client.NewRdsMySQLGormClient()
	assert.NoError(t, err)

	instance, err := clientInstance.GetInstance()
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestNewRdsMySQLGormClient_WithMockAuth_Fail(t *testing.T) {
	resetSingleton()
	clientInstance, err := client.NewRdsMySQLGormClient()
	assert.NoError(t, err)

	instance, err := clientInstance.GetInstance()
	assert.Nil(t, instance)
	assert.NotNil(t, err)
}
func TestNewRdsMySQLGormClient_ConfigSuccess(t *testing.T) {
	resetSingleton()

	clientInstance, err := client.NewRdsMySQLGormClient()
	assert.NoError(t, err)
	assert.NotNil(t, clientInstance)
}
