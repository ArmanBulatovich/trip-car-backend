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

func CreateOrganization(c *gin.Context) {
	admin, ok := getAdminFromContext(c)
	if !ok {
		log.Printf("admin_handlers.CreateOrganization->getAdminFromContext: Error\n")
		c.JSON(http.StatusBadGateway, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	req := &dto.OrganizationRequest{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.CreateOrganization->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.CreateOrganization(req, admin)
	if err != nil {
		log.Printf("admin_handlers.CreateOrganization->CreateOrganization: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateOrganization(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Printf("admin_handlers.UpdateOrganization->Atoi: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidUrlParam))
		return
	}

	req := &dto.OrganizationRequest{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.UpdateOrganization->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.UpdateOrganization(req, uint(id))
	if err != nil {
		log.Printf("admin_handlers.UpdateOrganization->UpdateOrganization: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetOrganizations(c *gin.Context) {
	req := &dto.GetOrganizationsRequest{}
	if err := c.BindQuery(&req); err != nil {
		log.Printf("admin_handlers.GetOrganizations->BindQuery: %s\n", err)
		c.JSON(http.StatusBadGateway, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	resp, err := admin_services.GetOrganizations(req)
	if err != nil {
		log.Printf("admin_handlers.GetOrganizations->GetOrganizations: %s\n", err)
		c.JSON(http.StatusBadGateway, responses.CreateErrorResponse(nil, "", responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}
