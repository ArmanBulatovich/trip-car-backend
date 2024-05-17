package repositories

import (
	"encoding/json"
	"log"

	"github.com/trip/trip-service/internal/db"
	"github.com/trip/trip-service/internal/models"
	"github.com/trip/trip-service/internal/utils"
)

func CreateOrganization(name, slug string, metadata []byte, adminId uint) (uint, error) {
	query := `
		INSERT INTO organizations(name, slug, metadata, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	var id uint
	err := db.DB.QueryRow(query, name, slug, metadata, adminId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetOrganizations(page, perPage int) ([]*models.Organization, error) {
	query := `
		SELECT id, name, slug, metadata
		FROM organizations
		WHERE deleted_at IS NULL
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	limit, offset := utils.GetLimitAndOffset(page, perPage)

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	organizations := make([]*models.Organization, 0)
	var metadata []byte

	for rows.Next() {
		organization := models.Organization{}

		rows.Scan(
			&organization.ID,
			&organization.Name,
			&organization.Slug,
			&metadata,
		)

		if err := json.Unmarshal(metadata, &organization.Metadata); err != nil {
			log.Printf("Organization[%d] metadata Unmarshal error: %s\n", organization.ID, err)
		}

		organizations = append(organizations, &organization)
	}

	return organizations, nil
}
