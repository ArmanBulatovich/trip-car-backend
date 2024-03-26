package dto

type CreateOrgUserRequestDTO struct {
	OrganizationID uint   `json:"organization_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}

type CreateOrgUserResponseDTO struct {
	ID uint `json:"id"`
}
