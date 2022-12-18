package config

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvLoad() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println("Error loading ./config/.env file")
	}
}
