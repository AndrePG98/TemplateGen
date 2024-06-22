package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DRIVER   string
	USER     string
	PASSWORD string
	HOST     string
	DB_PORT  string
	DB       string
	KEY      string
	PORT     string
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file: %w", err)
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func LoadConfig() *Config {
	loadEnv()

	return &Config{
		DRIVER:   getEnv("DRIVER"),
		USER:     getEnv("USER"),
		PASSWORD: getEnv("PASSWORD"),
		HOST:     getEnv("HOST"),
		DB_PORT:  getEnv("DB_PORT"),
		DB:       getEnv("DB"),
		KEY:      getEnv("KEY"),
		PORT:     getEnv("PORT"),
	}
}
