package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DBPort     string
	DBUser     string
	DBPassword string
	Port       string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")

	}

	config := &Config{
		DBName:     getEnvOrDefault("DB_NAME", "pokedot_db"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrDefault("DB_USER", "user"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "password"),
		Port:       getEnvOrDefault("PORT ", "3030"),
	}

	return config

}

// getEnvOrDefault returns the environment variable or the default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
