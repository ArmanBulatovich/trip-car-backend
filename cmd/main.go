package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/trip/trip-service/internal/api/server"
	"github.com/trip/trip-service/internal/config"
)

func main() {
	log.Println("Starting trip car backend...")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf(".env load error: %s\n" + err.Error())
	}

	err = config.Config()
	defer config.ConfigDefer()
	if err != nil {
		log.Println(err)
		return
	}

	server.Engine()
}
