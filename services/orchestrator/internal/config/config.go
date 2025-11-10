package config

import (
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Services ServicesConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type ServicesConfig struct {
	UserService     ServiceEndpoint
	TemplateService ServiceEndpoint
	UseMockServices bool
}

type ServiceEndpoint struct {
	BaseURL string
	Timeout time.Duration
}

type LoggingConfig struct {
	Level  string
	Format string // json or console
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 10*time.Second),
		},
		Services: ServicesConfig{
			UserService: ServiceEndpoint{
				BaseURL: getEnv("USER_SERVICE_URL", "http://user-service:8081"),
				Timeout: getDurationEnv("USER_SERVICE_TIMEOUT", 3*time.Second),
			},
			TemplateService: ServiceEndpoint{
				BaseURL: getEnv("TEMPLATE_SERVICE_URL", "http://template-service:8082"),
				Timeout: getDurationEnv("TEMPLATE_SERVICE_TIMEOUT", 3*time.Second),
			},
			UseMockServices: getBoolEnv("USE_MOCK_SERVICES", true),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true" || value == "1"
	}
	return defaultValue
}
