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

func CreateOrgUser(req *dto.CreateOrgUserRequestDTO, admin *models.Admin) (*responses.ApiResponse, error) {
	if !constants.IsCorrectOrgUserRole(req.Role) {
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

	id, err := repositories.CreateOrgUser(req.Email, hPassword, req.Role, req.OrganizationID, admin.ID)
	if err != nil {
		return nil, err
	}

	respDto := dto.CreateOrgUserResponseDTO{ID: id}
	resp := responses.CreateSuccessResponse(respDto, "ok")
	return &resp, nil
}

func GetOrgUsers(id uint) (*responses.ApiResponse, error) {
	users, err := repositories.GetOrgUsers(id)
	if err != nil {
		return nil, err
	}

	respDto := dto.GetOrgUsersResponseDTO{Users: users}
	resp := responses.CreateSuccessResponse(respDto, "")

	return &resp, nil
}
