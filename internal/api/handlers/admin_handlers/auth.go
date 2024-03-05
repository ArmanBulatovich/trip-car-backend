package admin_handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
)

func Login(c *gin.Context) {
	req := dto.LoginRequest{}

	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.LoginHandler->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, ""))
		return
	}

	resp, err := 
}
