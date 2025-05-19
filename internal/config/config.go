package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Port           string
	PostgresHost   string
	PostgresPort   string
	PostgresUser   string
	PostgresPass   string
	PostgresDBName string
	PostgresSSL    string
	PostgresURL    string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		Port:           getEnv("PORT", "8080"),
		PostgresHost:   getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:   getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:   getEnv("POSTGRES_USER", "postgres"),
		PostgresPass:   getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDBName: getEnv("POSTGRES_DB", "urlshortener"),
		PostgresSSL:    getEnv("POSTGRES_SSL_MODE", "disable"),
		PostgresURL:    getEnv("POSTGRES_URL", ""),
	}

	return config, nil
}

// PostgresDSN returns the PostgreSQL connection string
func (c *Config) PostgresDSN() string {
	if c.PostgresURL != "" {
		return c.PostgresURL
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresDBName,
		c.PostgresSSL,
	)
}

// getEnv reads an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
