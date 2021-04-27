package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}
	log.Printf("###%v\n", os.Getenv("TOKEN_SECRET"))
}
