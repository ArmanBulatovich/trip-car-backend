package repositories

import "github.com/trip/trip-service/internal/db"

func CreateOrgUser(email, password, role string, organizationID, adminID uint) (uint, error) {
	query := `
		INSERT INTO organization_users(email, password, role, organization_id, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	var id uint
	err := db.DB.QueryRow(query, email, password, role, organizationID, adminID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
