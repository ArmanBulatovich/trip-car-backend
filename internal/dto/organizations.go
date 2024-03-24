package dto

type CreateOrganizationRequestDTO struct {
	Name     string      `json:"name"`
	Slug     string      `json:"slug"`
	Metadata interface{} `json:"metadata"`
}

type CreateOrganizationResponseDTO struct {
	ID uint `json:"id"`
}
