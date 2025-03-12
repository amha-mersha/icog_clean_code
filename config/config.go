package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	APIVersion       string
	Port             int
}

func LoadConfig(dir string) *Config {
	if err := godotenv.Load(dir); err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Error loading .env file: Can't load PORT")
	}
	return &Config{
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "yourpassword"),
		PostgresDB:       getEnv("POSTGRES_DB", "task_manager"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		APIVersion:       getEnv("APIVersion", "v1"),
		Port:             port,
	}
}

func getEnv(key, fallback string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	return value
}
