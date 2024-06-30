package repositories

import (
	"log"

	"github.com/trip/trip-service/internal/db"
	"github.com/trip/trip-service/internal/models"
)

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

func GetOrgUsers(organizationID uint) ([]*models.OrgUser, error) {
	query := `
		SELECT id, email, role
		FROM organization_users
		WHERE organization_id = $1 AND deleted_at IS NULL
	`
	rows, err := db.DB.Query(query, organizationID)
	if err != nil {
		return nil, err
	}

	users := make([]*models.OrgUser, 0)

	for rows.Next() {
		user := models.OrgUser{}

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Role,
		)
		if err != nil {
			log.Printf("GetOrgUsers->scan error: organization[%d] error[%s]", organizationID, err.Error())
		}

		users = append(users, &user)
	}

	return users, nil
}
