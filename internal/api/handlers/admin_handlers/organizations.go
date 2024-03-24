package admin_handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/services/admin_services"
)

func CreateOrganization(c *gin.Context) {
	req := &dto.CreateOrganizationRequestDTO{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.CreateOrganization->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.CreateOrganization(req)
	if err != nil {
		log.Printf("admin_handlers.CreateOrganization->CreateOrganization: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}