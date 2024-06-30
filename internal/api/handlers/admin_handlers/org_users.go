package admin_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/services/admin_services"
)

func CreateOrgUser(c *gin.Context) {
	admin, ok := getAdminFromContext(c)
	if !ok {
		log.Printf("admin_handlers.CreateOrgUser->getAdminFromContext: Error\n")
		c.JSON(http.StatusBadGateway, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	req := &dto.CreateOrgUserRequestDTO{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.CreateOrgUser->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.CreateOrgUser(req, admin)
	if err != nil {
		log.Printf("admin_handlers.CreateOrgUser->CreateOrgUser: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetOrgUsers(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Printf("admin_handlers.GetOrgUsers->ParseUint: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidUrlParam", responses.InvalidUrlParam))
		return
	}

	resp, err := admin_services.GetOrgUsers(uint(id))
	if err != nil {
		log.Printf("admin_handlers.GetOrgUsers->GetOrgUsers: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}
