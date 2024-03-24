package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/handlers/admin_handlers"
)

func SetAdminRoutesV1(engine *gin.Engine) {
	v1 := engine.Group("/api/admin/v1/")

	{
		v1.POST("/login", admin_handlers.Login)
	}
	{
		v1.POST("/organizations", admin_handlers.CreateOrganization)
	}
}
