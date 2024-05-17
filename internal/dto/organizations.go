package dto

import "github.com/trip/trip-service/internal/models"

type OrganizationRequest struct {
	Name     string      `json:"name"`
	Slug     string      `json:"slug"`
	Metadata interface{} `json:"metadata"`
}

type OrganizationResponse struct {
	ID uint `json:"id"`
}

type GetOrganizationsRequest struct {
	Pagination
}

type GetOrganizationsResponse struct {
	Organizations []*models.Organization `json:"organizations"`
}

type GetOrganizationResponse struct {
	Organization *models.Organization `json:"organization"`
}
