package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
}

var AppConfig *Config

func LoadConfig() {

	envFile := "/Users/khushi.nijhawan/GolandProjects/HackOn/.env"
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	AppConfig = &Config{
		DBUser:     os.Getenv("DATABASE_USER"),
		DBPassword: os.Getenv("DATABASE_PASSWORD"),
		DBName:     os.Getenv("DATABASE_NAME"),
		DBPort:     os.Getenv("DATABASE_PORT"),
		DBHost:     os.Getenv("DATABASE_HOST"),
	}

	if AppConfig.DBUser == "" || AppConfig.DBPassword == "" || AppConfig.DBName == "" || AppConfig.DBPort == "" || AppConfig.DBHost == "" {
		log.Fatalf("One or more required environment variables are missing")
	}
}
