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
		c.JSON(http.StatusUnauthorized, responses.CreateErrorResponse(nil, "InvalidAuthToken", responses.InvalidAuthToken))
		return
	}

	req := &dto.OrganizationRequest{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.CreateOrganization->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidRequestBody", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.CreateOrganization(req, admin)
	if err != nil {
		log.Printf("admin_handlers.CreateOrganization->CreateOrganization: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, err.Error(), responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateOrganization(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Printf("admin_handlers.UpdateOrganization->Atoi: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidUrlParam", responses.InvalidUrlParam))
		return
	}

	req := &dto.OrganizationRequest{}
	if err := c.BindJSON(req); err != nil {
		log.Printf("admin_handlers.UpdateOrganization->BindJSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidRequestBody", responses.InvalidRequestBody))
		return
	}

	resp, err := admin_services.UpdateOrganization(req, uint(id))
	if err != nil {
		log.Printf("admin_handlers.UpdateOrganization->UpdateOrganization: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, err.Error(), responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetOrganization(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Printf("admin_handlers.GetOrganization->ParseUint: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidUrlParam", responses.InvalidUrlParam))
		return
	}

	resp, err := admin_services.GetOrganization(uint(id))
	if err != nil {
		log.Printf("admin_handlers.GetOrganization->GetOrganization: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, err.Error(), responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteOrganization(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Printf("admin_handlers.DeleteOrganization->ParseUint: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidUrlParam", responses.InvalidUrlParam))
		return
	}

	resp, err := admin_services.DeleteOrganization(uint(id))
	if err != nil {
		log.Printf("admin_handlers.DeleteOrganization->DeleteOrganization: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, err.Error(), responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetOrganizations(c *gin.Context) {
	req := &dto.GetOrganizationsRequest{}
	if err := c.BindQuery(&req); err != nil {
		log.Printf("admin_handlers.GetOrganizations->BindQuery: %s\n", err)
		c.JSON(http.StatusBadRequest, responses.CreateErrorResponse(nil, "InvalidUrlParams", responses.InvalidUrlParam))
		return
	}

	resp, err := admin_services.GetOrganizations(req)
	if err != nil {
		log.Printf("admin_handlers.GetOrganizations->GetOrganizations: %s\n", err)
		c.JSON(http.StatusInternalServerError, responses.CreateErrorResponse(nil, err.Error(), responses.Error))
		return
	}

	c.JSON(http.StatusOK, resp)
}
