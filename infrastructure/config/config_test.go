package config_test

import (
	"aws_challenge_pragma/infrastructure/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_Success(t *testing.T) {
	tempEnv := ".env"
	content := "TEST_ENV=test_value"
	os.WriteFile(tempEnv, []byte(content), 0644)
	defer os.Remove(tempEnv)

	result := config.LoadConfig()
	assert.True(t, result, "LoadConfig should return true when .env is present")
}

func TestLoadConfig_Failure(t *testing.T) {
	os.Remove(".env")

	result := config.LoadConfig()
	assert.False(t, result, "LoadConfig should return false when .env is missing")
}

func TestGetEnv_ExistingVariable(t *testing.T) {
	os.Setenv("TEST_ENV", "test_value")
	defer os.Unsetenv("TEST_ENV")

	result := config.GetEnv("TEST_ENV")
	assert.Equal(t, "test_value", result, "GetEnv should return the existing environment variable")
}

func TestGetEnv_DefaultValue(t *testing.T) {
	result := config.GetEnv("NON_EXISTENT_ENV", "default_value")
	assert.Equal(t, "default_value", result, "GetEnv should return the default value when the environment variable is missing")
}

func TestGetEnv_EmptyWhenMissing(t *testing.T) {
	result := config.GetEnv("NON_EXISTENT_ENV")
	assert.Equal(t, "", result, "GetEnv should return an empty string when the environment variable is missing and no default is provided")
}
