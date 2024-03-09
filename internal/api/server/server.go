package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/routers"
)

func Engine() {
	engine := gin.Default()

	routers.SetAdminRoutesV1(engine)

	err := engine.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
	if err != nil {
		log.Printf("GIN: engine.Run error: %s\n", err.Error())
	}
}
