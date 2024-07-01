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

func UpdateOrganization(id uint, name, slug string, metadata []byte) error {
	query := `
		UPDATE organizations
		SET name = $2, slug = $3, metadata = $4
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := db.DB.Exec(query, id, name, slug, metadata)
	return err
}

func GetOrganization(id uint) (*models.Organization, error) {
	query := `
		SELECT name, slug, metadata
		FROM organizations
		WHERE id = $1 AND deleted_at IS NULL
	`

	organization := models.Organization{}

	var metadata []byte
	err := db.DB.QueryRow(query, id).Scan(&organization.Name, &organization.Slug, &metadata)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(metadata, &organization.Metadata); err != nil {
		log.Printf("Organization[%d] metadata Unmarshal error: %s\n", organization.ID, err)
	}

	organization.ID = id

	return &organization, nil
}

func DeleteOrganization(id uint) error {
	query := `
		DELETE FROM organizations
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := db.DB.Exec(query, id)
	return err
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
