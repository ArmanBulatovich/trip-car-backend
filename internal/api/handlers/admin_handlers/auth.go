package admin_handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/services/admin_services"
)

func Login(c *gin.Context) {
	req := &dto.LoginRequestDTO{}

	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.LoginHandler->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.Login(req)
	if err != nil {
		log.Printf("admin_handlers.LoginHandler->Login: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	c.JSON(http.StatusOK, responses.CreateSuccessResponse(resp, "ok"))
}
