package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	_ = godotenv.Load()
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
