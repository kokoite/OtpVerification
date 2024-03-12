package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvAccountsSID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func getEnvAuthToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error while loading auth token")
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func getEnvServiceID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error while loading env file")
	}
	return os.Getenv("TWILIO_SERVICE_ID")
}
