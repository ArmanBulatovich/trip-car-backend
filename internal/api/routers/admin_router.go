package routers

import "github.com/gin-gonic/gin"

func SetAdminRoutesV1(engine *gin.Engine) {
	v1 := engine.Group("/api/admin/v1/")

	{
		v1.POST("/login")
	}
}
