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

	port := getEnv("PORT", "")
	if port == "" {
		port = getEnv("APP_PORT", "3000")
	}
	if len(port) > 0 && port[0] != ':' {
		port = ":" + port
	}

	return AppConfig{
		Port:     port,
		DBDriver: getEnv("DB_DRIVER", "mysql"),
		DBSource: getEnv("DB_SOURCE", ""),
	}
}

// getEnv is a simple helper function to read an environment variable or return a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
