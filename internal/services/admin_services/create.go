package admin_services

import (
	"fmt"

	"github.com/trip/trip-service/internal/constants"
	"github.com/trip/trip-service/internal/repositories"
	"github.com/trip/trip-service/internal/utils"
)

func CreateSuperAdmin(email, password string) error {
	hPassword := utils.HashSha256(password)

	exists, err := repositories.IsAdminExists(email)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("admin [%s] exists", email)
	}

	return repositories.CreateAdmin(email, hPassword, constants.AdminTypeSuperAdmin)
}
