package dto

import "github.com/trip/trip-service/internal/models"

type CreateOrganizationRequestDTO struct {
	Name     string      `json:"name"`
	Slug     string      `json:"slug"`
	Metadata interface{} `json:"metadata"`
}

type CreateOrganizationResponseDTO struct {
	ID uint `json:"id"`
}

type GetOrganizationsRequest struct {
	Pagination
}

type GetOrganizationsResponse struct {
	Organizations []*models.Organization `json:"organizations"`
}
