package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type LoggerLevel string

const (
	INFO    LoggerLevel = "INFO"
	ERROR   LoggerLevel = "ERROR"
	WARNING LoggerLevel = "WARNING"
)

type LogEntry struct {
	Timestamp   string                 `json:"timestamp"`
	ID          string                 `json:"id"`
	Code        string                 `json:"code"`
	Message     string                 `json:"message"`
	Detail      string                 `json:"detail"`
	Payload     map[string]interface{} `json:"payload"`
	Level       LoggerLevel            `json:"level"`
	Severity    LoggerLevel            `json:"severity"`
	Service     string                 `json:"service"`
	Environment string                 `json:"environment"`
}

type Logger struct {
	id          string
	code        string
	message     string
	detail      string
	metadata    map[string]interface{}
	level       LoggerLevel
	service     string
	environment string
}

func NewLogger(level LoggerLevel) *Logger {
	return &Logger{
		level:       level,
		service:     GetEnv("SERVICE_NAME", "default-service"),
		environment: GetEnv("NODE_ENV", "development"),
		metadata:    make(map[string]interface{}),
	}
}

func (l *Logger) SetCode(code string) *Logger {
	l.code = code
	return l
}

func (l *Logger) SetMessage(message string) *Logger {
	l.message = message
	return l
}

func (l *Logger) SetDetail(detail string) *Logger {
	l.detail = detail
	return l
}

func (l *Logger) SetMetadata(metadata map[string]interface{}) *Logger {
	l.metadata = metadata
	return l
}

func (l *Logger) Write() {
	if l.code == "" || l.message == "" || l.detail == "" {
		return
	}

	l.id = uuid.New().String()

	logEntry := LogEntry{
		Timestamp:   time.Now().Format(time.RFC3339),
		ID:          l.id,
		Code:        l.code,
		Message:     l.message,
		Detail:      l.detail,
		Payload:     l.metadata,
		Level:       l.level,
		Severity:    l.level,
		Service:     l.service,
		Environment: l.environment,
	}

	logJSON, _ := json.Marshal(logEntry)
	fmt.Println(string(logJSON))

	l.reset()
}

func (l *Logger) reset() {
	l.id = ""
	l.code = ""
	l.message = ""
	l.detail = ""
	l.metadata = make(map[string]interface{})
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

type AppLogger struct{}

func (a AppLogger) Info() *Logger {
	return NewLogger(INFO)
}

func (a AppLogger) Warn() *Logger {
	return NewLogger(WARNING)
}

func (a AppLogger) Error() *Logger {
	return NewLogger(ERROR)
}
