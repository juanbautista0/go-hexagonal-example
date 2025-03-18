package config

import (
	"aws_challenge_pragma/infrastructure/logger"
	"os"

	"github.com/joho/godotenv"
)

var LoadConfig = loadConfig
var GetEnv = getEnv

func loadConfig() bool {
	err := godotenv.Load()

	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "prod" {
		return true
	}
	ensureEnvVariable("_LAMBDA_SERVER_PORT", "8080")
	ensureEnvVariable("AWS_LAMBDA_RUNTIME_API", "localhost")
	if err != nil {
		appLogger := logger.AppLogger{}

		appLogger.Error().
			SetCode("CONFIG_001").
			SetDetail("Failed to load environment variables from .env file").
			SetMessage("Configuration loading failed").
			SetMetadata(map[string]interface{}{
				"error": err.Error(),
			}).
			Write()

		return false
	}

	return true
}

func getEnv(key string, defaultValue ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	switch len(defaultValue) {
	case 1:
		return defaultValue[0]
	default:
		return ""
	}
}

func ensureEnvVariable(key, defaultValue string) {
	if _, exists := os.LookupEnv(key); !exists {
		os.Setenv(key, defaultValue)
	}
}
