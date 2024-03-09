package repositories

import (
	"github.com/trip/trip-service/internal/db"
	"github.com/trip/trip-service/internal/models"
)

func GetAdmin(email, password string) (*models.Admin, error) {
	query := `
		SELECT id
		FROM admins
		WHERE email = $1 AND password = $2 AND deleted_at IS NULL
	`
	admin := models.Admin{}

	row := db.DB.QueryRow(query, email, password)

	err := row.Scan(&admin.ID)
	if err != nil {
		return nil, err
	}

	admin.Email = email
	return &admin, nil
}
