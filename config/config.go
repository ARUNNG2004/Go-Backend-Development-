package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds all the environment variables
type AppConfig struct {
	Port     string
	DBDriver string
	DBSource string
}

// LoadConfig reads the .env file and populates the AppConfig struct
func LoadConfig() AppConfig {
	// Load the .env file. If it fails, we just log a warning
	// (it might fail in production where env vars are injected directly by Docker/Kubernetes)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	return AppConfig{
		Port:     getEnv("APP_PORT", ":3000"), // Defaults to :3000 if not found
		DBDriver: getEnv("DB_DRIVER", "mysql"),
		DBSource: getEnv("DB_SOURCE", ""), // We leave this empty so it errors out if missing
	}
}

// getEnv is a simple helper function to read an environment variable or return a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
