package repositories

import (
	"github.com/trip/trip-service/internal/db"
	"github.com/trip/trip-service/internal/models"
)

func CreateProvUser(user *models.ProvUser, adminID uint) (uint, error) {
	query := `
		INSERT INTO provider_users(email, password, role, phone_number, fullname, provider_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`
	var id uint
	err := db.DB.QueryRow(query, user.Email, user.Password, user.Role, user.PhoneNumber, user.Fullname, user.ProviderID, adminID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
