package models

type OrgUser struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}
