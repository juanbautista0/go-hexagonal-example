package logger_test

import (
	"aws_challenge_pragma/infrastructure/logger"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger_NotNil(t *testing.T) {
	log := logger.NewLogger(logger.INFO)
	assert.NotNil(t, log, "Logger instance should not be nil")
}

func TestLogger_Write_PrintsLog(t *testing.T) {
	// Capturar la salida estándar
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	log := logger.NewLogger(logger.ERROR).
		SetCode("ERR001").
		SetMessage("Test message").
		SetDetail("Test detail").
		SetMetadata(map[string]interface{}{"key": "value"})

	log.Write()

	// Restaurar salida estándar
	w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)

	assert.Contains(t, buf.String(), `"code":"ERR001"`)
	assert.Contains(t, buf.String(), `"message":"Test message"`)
	assert.Contains(t, buf.String(), `"detail":"Test detail"`)
	assert.Contains(t, buf.String(), `"level":"ERROR"`)
}

func TestLogger_Write_ResetsValues(t *testing.T) {
	log := logger.NewLogger(logger.WARNING)
	log.SetCode("WARN001").
		SetMessage("Test warning").
		SetDetail("Test detail")

	log.Write()

	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	log.Write()

	w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)

	assert.Equal(t, "", buf.String())
}

func TestAppLogger_Info(t *testing.T) {
	appLogger := logger.AppLogger{}
	log := appLogger.Info()
	assert.NotNil(t, log)
}

func TestAppLogger_Warn(t *testing.T) {
	appLogger := logger.AppLogger{}
	log := appLogger.Warn()
	assert.NotNil(t, log)
}

func TestAppLogger_Error(t *testing.T) {
	appLogger := logger.AppLogger{}
	log := appLogger.Error()
	assert.NotNil(t, log)
}

func TestLogger_Reset(t *testing.T) {
	log := logger.NewLogger(logger.INFO).
		SetCode("RESET001").
		SetMessage("Reset test").
		SetDetail("Testing reset").
		SetMetadata(map[string]interface{}{"key": "value"})

	log.Write()

	assert.NotNil(t, log)

}

func TestGetEnv_Default(t *testing.T) {
	defaultValue := "default-service"
	value := logger.GetEnv("NON_EXISTENT_ENV", defaultValue)
	assert.Equal(t, defaultValue, value)
}

func TestGetEnv_SetValue(t *testing.T) {
	os.Setenv("TEST_ENV", "custom-value")
	defer os.Unsetenv("TEST_ENV")

	value := logger.GetEnv("TEST_ENV", "default-value")
	assert.Equal(t, "custom-value", value)
}
