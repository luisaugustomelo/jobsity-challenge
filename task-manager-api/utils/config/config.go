package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}

func LoadEnv(path ...string) {
	envPath := ".env"
	if path != nil {
		envPath = path[0]
	}
	err := godotenv.Overload(envPath)
	if err != nil {
		log.Println("unable to load .env, running with default env")
	}

	PORT = GetEnv("PORT", "3333")
}
