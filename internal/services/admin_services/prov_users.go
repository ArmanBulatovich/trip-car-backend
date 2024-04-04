package admin_services

import (
	"fmt"

	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/constants"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/models"
	"github.com/trip/trip-service/internal/repositories"
	"github.com/trip/trip-service/internal/utils"
)

func CreateProvUser(req *dto.CreateProvUserRequestDTO, admin *models.Admin) (*responses.ApiResponse, error) {
	if !constants.IsCorrectProvUserRole(req.Role) {
		resp := responses.CreateErrorResponse(nil, "invalid role", responses.InvalidRole)
		return &resp, nil
	}
	if !utils.IsValidEmail(req.Email) {
		resp := responses.CreateErrorResponse(nil, "invalid email", responses.InvalidEmail)
		return &resp, nil
	}

	password, err := utils.GeneratePassword()
	if err != nil {
		resp := responses.CreateErrorResponse(nil, "gen password error", responses.Error)
		return &resp, nil
	}
	hPassword := utils.HashSha256(password)

	// TODO: SEND EMAIL
	fmt.Println(password)

	user := models.ProvUser{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Fullname:    req.Fullname,
		Password:    hPassword,
		Role:        req.Role,
		ProviderID:  req.ProviderID,
	}

	id, err := repositories.CreateProvUser(&user, admin.ID)
	if err != nil {
		return nil, err
	}

	respDto := dto.CreateOrgUserResponseDTO{ID: id}
	resp := responses.CreateSuccessResponse(respDto, "ok")
	return &resp, nil
}
