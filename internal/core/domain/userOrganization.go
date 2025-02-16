package domain

type UserOrganization struct {
	ID             string `db:"organizacion_por_usuario_id" json:"id"`
	OrganizationID string `db:"organizacion_id" json:"organization_id"`
	UserID         string `db:"usuario_id" json:"user_id"`

	// This is a pointer to the User struct
	User         *User
	Organization *Organization
}
