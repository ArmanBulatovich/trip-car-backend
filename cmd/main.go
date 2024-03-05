package main

import (
	"log"

	"github.com/trip/trip-service/internal/api/server"
)

func main() {
	log.Println("Starting trip car backend...")

	server.Engine()
}
