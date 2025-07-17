package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoURI      string
	Port          string
	SessionSecret string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		log.Fatal("MONGO_URI must be set")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	SessionSecret = os.Getenv("SESSION_SECRET")

	if SessionSecret == "" {
		SessionSecret = "some_secret"
	}
}
