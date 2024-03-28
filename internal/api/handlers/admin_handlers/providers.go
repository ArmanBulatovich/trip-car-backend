package admin_handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/services/admin_services"
)

func CreateProvider(c *gin.Context) {
	admin, ok := getAdminFromContext(c)
	if !ok {
		log.Printf("admin_handlers.CreateProvider->getAdminFromContext: Error\n")
		c.JSON(http.StatusBadGateway, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	req := &dto.CreateProviderRequestDTO{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.CreateProvider->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.CreateProvider(req, admin)
	if err != nil {
		log.Printf("admin_handlers.CreateProvider->CreateProvider: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}
