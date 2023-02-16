package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Token() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("Не удалось найти .env файл!!!")
	}

	token := os.Getenv("TOKEN")
	return token
}
