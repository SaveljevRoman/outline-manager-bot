package fileLoader

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvLoader() {
	var envLoad []string
	envLoad = append(envLoad, ".env")
	if err := godotenv.Load(envLoad...); err != nil {
		log.Fatal("No .env file found")
	}
}
