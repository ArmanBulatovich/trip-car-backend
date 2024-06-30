package dto

import "github.com/trip/trip-service/internal/models"

type CreateOrgUserRequestDTO struct {
	OrganizationID uint   `json:"organization_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}

type CreateOrgUserResponseDTO struct {
	ID uint `json:"id"`
}

type GetOrgUsersResponseDTO struct {
	Users []*models.OrgUser `json:"users"`
}
