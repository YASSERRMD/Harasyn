package http

import (
	"os"
	"strconv"
)

type Config struct {
	Port          int
	Host          string
	DatabaseURL   string
	RedisURL      string
	NATSURL       string
	JWTSecret     string
	CORSOrigins   []string
	RateLimit     int
	RateWindow    int
	LogLevel      string
	Environment   string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnvInt("PORT", 8080),
		Host:        getEnv("HOST", "0.0.0.0"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://harasyn:harasyn@localhost:5432/harasyn?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
		NATSURL:     getEnv("NATS_URL", "nats://localhost:4222"),
		JWTSecret:   getEnv("JWT_SECRET", "dev-secret-change-in-production"),
		CORSOrigins: []string{getEnv("CORS_ORIGINS", "http://localhost:3000")},
		RateLimit:   getEnvInt("RATE_LIMIT", 100),
		RateWindow:  getEnvInt("RATE_WINDOW", 60),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
