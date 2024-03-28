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

func CreateProvider(req *dto.CreateProviderRequestDTO, admin *models.Admin) (*responses.ApiResponse, error) {
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

	provider := models.Provider{
		Name:        req.Name,
		Slug:        req.Slug,
		Metadata:    metadataJson,
		BIN:         req.BIN,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}

	id, err := repositories.CreateProvider(&provider, admin.ID)
	if err != nil {
		return nil, err
	}

	respDto := dto.CreateProviderResponseDTO{ID: id}
	resp := responses.CreateSuccessResponse(respDto, "ok")
	return &resp, nil
}
