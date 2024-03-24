package admin_services

import (
	"encoding/json"
	"strings"

	"github.com/trip/trip-service/internal/api/responses"
	"github.com/trip/trip-service/internal/dto"
	"github.com/trip/trip-service/internal/repositories"
	"github.com/trip/trip-service/internal/utils"
)

func CreateOrganization(req *dto.CreateOrganizationRequestDTO) (*responses.ApiResponse, error) {
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

	id, err := repositories.CreateOrganization(req.Name, req.Slug, metadataJson, 1)
	if err != nil {
		return nil, err
	}

	respDto := dto.CreateOrganizationResponseDTO{ID: id}
	resp := responses.CreateSuccessResponse(respDto, "ok")
	return &resp, nil
}
