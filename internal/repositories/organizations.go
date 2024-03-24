package repositories

import "github.com/trip/trip-service/internal/db"

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
