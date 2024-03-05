package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/trip/trip-service/internal/auth"
	"github.com/trip/trip-service/internal/models"
)

func CreateAdminToken(admin *models.Admin) (string, error) {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // TODO time.HOUR * 24 -> const
	}
	claims := auth.TokenClaims{
		UserId:         admin.ID,
		UserEmail:      admin.Email,
		UserType:       auth.AdminUserType,
		StandardClaims: standardClaims,
	}
	return auth.GenerateToken(claims)
}
