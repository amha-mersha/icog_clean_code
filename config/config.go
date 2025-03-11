package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
}

func LoadConfig(dir string) *Config {
	if err := godotenv.Load(dir); err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "yourpassword"),
		PostgresDB:       getEnv("POSTGRES_DB", "task_manager"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	return value
}
