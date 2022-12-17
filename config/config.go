package config

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvLoad() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading ./config/.env file")
	}
}
