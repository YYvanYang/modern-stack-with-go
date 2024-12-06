package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxIdleConns    int           `default:"10"`
	MaxOpenConns    int           `default:"100"`
	ConnMaxLifetime time.Duration `default:"1h"`
}

type JWTConfig struct {
	Secret        string
	TokenDuration time.Duration
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			DBName:         getEnv("DB_NAME", "modern_stack"),
			SSLMode:        getEnv("DB_SSLMODE", "disable"),
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: time.Hour,
		},
		JWT: JWTConfig{
			Secret:        mustGetEnv("JWT_SECRET"),
			TokenDuration: time.Hour * 24,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func mustGetEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	if os.Getenv("GIN_MODE") == "debug" {
		switch key {
		case "JWT_SECRET":
			return "development_jwt_secret_key"
		}
	}
	log.Fatalf("Environment variable %s is required", key)
	return ""
} 