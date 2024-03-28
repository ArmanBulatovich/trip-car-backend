package repositories

import (
	"github.com/trip/trip-service/internal/db"
	"github.com/trip/trip-service/internal/models"
)

func CreateProvider(provider *models.Provider, adminID uint) (uint, error) {
	query := `
		INSERT INTO providers(name, slug, metadata, bin, email, phone_number, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	err := db.DB.QueryRow(
		query,
		provider.Name,
		provider.Slug,
		provider.Metadata,
		provider.BIN,
		provider.Email,
		provider.PhoneNumber,
		adminID).Scan(&provider.ID)
	if err != nil {
		return 0, err
	}

	return provider.ID, nil
}
