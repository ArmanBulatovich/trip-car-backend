package constants

const (
	AdminTypeSuperAdmin = "superadmin"
)

const (
	OrgUserRoleAdmin   = "admin"
	OrgUserRoleManager = "manager"
)

func IsCorrectOrgUserRole(role string) bool {
	return role == OrgUserRoleAdmin || role == OrgUserRoleManager
}
