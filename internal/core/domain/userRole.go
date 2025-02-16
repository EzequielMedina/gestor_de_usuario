package domain

type UserRole struct {
	ID     string `db:"usuario_por_rol_id" json:"id"`
	UserID string `db:"usuario_id" json:"user_id"`
	RoleID string `db:"rol_id" json:"role_id"`

	// Relations
	User *User
	Role *Role
}
