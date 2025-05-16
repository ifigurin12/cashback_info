package user

type User struct {
	ID       string   `json:"id" binding:"required"`
	Username string   `json:"username" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	RoleType RoleType `json:"role_type" binding:"required"`
	Phone    *string  `json:"phone,omitempty"`
}

type RoleType string

const (
	Default RoleType = "default"
	Admin   RoleType = "admin"
)

func GenerateRoleTypeFromString(value string) *RoleType {
	var result RoleType
	switch value {
	case "default":
		result = Default
	case "admin":
		result = Admin
	default:
		return nil
	}

	return &result
}
