package admin_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/models"
)

func getAdminFromContext(c *gin.Context) (*models.Admin, bool) {
	adm, ok := c.Get("admin")
	if !ok {
		return nil, false
	}
	admin, ok := adm.(models.Admin)
	if !ok {
		return nil, false
	}
	return &admin, true
}
