package admin_services

import (
	"fmt"

	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/repositories"
	"github.com/trip/trip-service/internal/utils"
)

func Login(req *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error) {
	hPassword := utils.HashSha256(req.Password)

	admin, err := repositories.GetAdmin(req.Email, hPassword)
	if err != nil {
		return nil, fmt.Errorf("GetAdmin: %s", err)
	}

	token, err := utils.CreateAdminToken(admin)
	if err != nil {
		return nil, fmt.Errorf("CreateAdminToken: %s", err)
	}

	return &dto.LoginResponseDTO{Token: token}, nil
}
