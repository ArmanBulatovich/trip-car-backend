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

func IsAdminExists(email string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM admins WHERE email = $1 AND deleted_at IS NULL
		);
	`

	var exists bool

	err := db.DB.QueryRow(query, email).Scan(&exists)

	return exists, err
}

func CreateAdmin(email, password, role string) error {
	query := `
		INSERT INTO admins
		(email, password, role)
		VALUES
		($1, $2, $3)
	`

	_, err := db.DB.Exec(query, email, password, role)

	return err
}
