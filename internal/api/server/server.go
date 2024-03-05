package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Engine() {
	engine := gin.Default()

	err := engine.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
	if err != nil {
		log.Printf("GIN: engine.Run error: %s\n", err.Error())
	}
}
