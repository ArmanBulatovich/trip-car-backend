package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/routers"
)

func Engine() {
	engine := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"}, // Укажите здесь домен вашего фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	engine.Use(cors.New(config))

	routers.SetAdminRoutesV1(engine)

	err := engine.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
	if err != nil {
		log.Printf("GIN: engine.Run error: %s\n", err.Error())
	}
}
