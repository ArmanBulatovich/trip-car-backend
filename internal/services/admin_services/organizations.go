package admin_services

import (
	"encoding/json"
	"strings"

	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/models"
	"github.com/trip/trip-service/internal/repositories"
	"github.com/trip/trip-service/internal/utils"
)

func CreateOrganization(req *dto.OrganizationRequest, admin *models.Admin) (*responses.ApiResponse, error) {
	req.Slug = strings.ToLower(req.Slug)

	if req.Name == "" {
		resp := responses.CreateErrorResponse(nil, "invalid name", responses.InvalidName)
		return &resp, nil
	}
	if !utils.IsValidSlug(req.Slug) {
		resp := responses.CreateErrorResponse(nil, "invalid slug", responses.InvalidSlug)
		return &resp, nil
	}

	metadataJson, err := json.Marshal(req.Metadata)
	if err != nil {
		resp := responses.CreateErrorResponse(nil, "invalid metadata", responses.InvalidMetadata)
		return &resp, nil
	}

	id, err := repositories.CreateOrganization(req.Name, req.Slug, metadataJson, admin.ID)
	if err != nil {
		return nil, err
	}

	respDto := dto.OrganizationResponse{ID: id}
	resp := responses.CreateSuccessResponse(respDto, "ok")
	return &resp, nil
}

func UpdateOrganization(req *dto.OrganizationRequest, id uint) (*responses.ApiResponse, error) {
	req.Slug = strings.ToLower(req.Slug)

	if req.Name == "" {
		resp := responses.CreateErrorResponse(nil, "invalid name", responses.InvalidName)
		return &resp, nil
	}
	if !utils.IsValidSlug(req.Slug) {
		resp := responses.CreateErrorResponse(nil, "invalid slug", responses.InvalidSlug)
		return &resp, nil
	}

	metadataJson, err := json.Marshal(req.Metadata)
	if err != nil {
		resp := responses.CreateErrorResponse(nil, "invalid metadata", responses.InvalidMetadata)
		return &resp, nil
	}

	err = repositories.UpdateOrganization(id, req.Name, req.Slug, metadataJson)
	if err != nil {
		return nil, err
	}

	resp := responses.CreateSuccessResponse(nil, "")
	return &resp, nil
}

func GetOrganizations(req *dto.GetOrganizationsRequest) (*responses.ApiResponse, error) {
	organizations, err := repositories.GetOrganizations(req.Page, req.PerPage)
	if err != nil {
		return nil, err
	}

	respDto := dto.GetOrganizationsResponse{
		Organizations: organizations,
	}
	resp := responses.CreateSuccessResponse(respDto, "ok")
	return &resp, nil
}
