package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"kennkitchen/config"
	"os"
)

func main() {
	file, err := os.OpenFile("logs/webserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.RunServer()
}
