package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/handlers/admin_handlers"
	"github.com/trip/trip-service/internal/api/middlewares"
)

func SetAdminRoutesV1(engine *gin.Engine) {
	v1 := engine.Group("/api/admin/v1/")

	{
		v1.POST("/login", admin_handlers.Login)
		v1.GET("/me", middlewares.AdminAuthMiddleware(), admin_handlers.Me)
	}
	{
		v1.POST("/organizations", middlewares.AdminAuthMiddleware(), admin_handlers.CreateOrganization)
		v1.PUT("/organizations/:id", middlewares.AdminAuthMiddleware(), admin_handlers.UpdateOrganization)
		v1.GET("/organizations/:id", middlewares.AdminAuthMiddleware(), admin_handlers.GetOrganization)
		v1.GET("/organizations", middlewares.AdminAuthMiddleware(), admin_handlers.GetOrganizations)
	}
	{
		v1.POST("/organizations/users", middlewares.AdminAuthMiddleware(), admin_handlers.CreateOrgUser)
		v1.GET("/organizations/:id/users", middlewares.AdminAuthMiddleware(), admin_handlers.GetOrgUsers)
	}
	{
		v1.POST("/providers", middlewares.AdminAuthMiddleware(), admin_handlers.CreateProvider)
		v1.POST("/providers/users", middlewares.AdminAuthMiddleware())
	}
}
