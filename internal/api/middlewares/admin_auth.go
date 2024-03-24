package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/auth"
	"github.com/trip/trip-service/internal/models"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 {
			resp := responses.CreateErrorResponse(nil, "invalid auth token", responses.InvalidAuthToken)
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		claims := auth.VerifyToken(tokenParts[1])
		if claims.ExpiresAt < time.Now().Unix() {
			resp := responses.CreateErrorResponse(nil, "token expired", responses.TokenExpired)
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		if claims.UserType != auth.AdminUserType {
			resp := responses.CreateErrorResponse(nil, "not admin", responses.Error)
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		admin := models.Admin{ID: claims.UserId, Email: claims.UserEmail}
		c.Set("admin", admin)
	}
}
