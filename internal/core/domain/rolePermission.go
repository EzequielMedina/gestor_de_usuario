package domain

type RolePermission struct {
	ID           string `db:"rol_por_permiso_id" json:"id"`
	RoleID       string `db:"rol_id" json:"role_id"`
	PermissionID string `db:"permiso_id" json:"permission_id"`

	Role       *Role
	Permission *Permission
}
