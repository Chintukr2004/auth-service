package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string

	AccessTokenExp string
	RefreshTokenExp string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "postgres"),
		DBUser:     getEnv("DB_USER", "postges"),
		DBName:     getEnv("DB_NAME", "auth_db"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
		AccessTokenExp: getEnv("ACCES_TOKEN_EXP","15m"),
		RefreshTokenExp: getEnv("REFRESH_TOKEN_EXP", "168h"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
