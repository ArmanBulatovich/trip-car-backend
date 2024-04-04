package constants

const (
	AdminTypeSuperAdmin = "superadmin"
)

const (
	OrgUserRoleAdmin   = "admin"
	OrgUserRoleManager = "manager"
)

const (
	ProvUserRoleAdmin   = "admin"
	ProvUserRoleManager = "manager"
)

func IsCorrectOrgUserRole(role string) bool {
	return role == OrgUserRoleAdmin || role == OrgUserRoleManager
}

func IsCorrectProvUserRole(role string) bool {
	return role == ProvUserRoleAdmin || role == ProvUserRoleManager
}
