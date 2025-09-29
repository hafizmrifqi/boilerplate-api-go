package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
	ServerPort string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env not found, using system env")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", ""),
		DBName:     getEnv("DB_NAME", "myapp_db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}