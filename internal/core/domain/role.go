package domain

type Role struct {
	ID          string  `db:"rol_id" json:"id"`
	Name        string  `db:"nombre" json:"name"`
	Description *string `db:"descripcion" json:"description,omitempty"`
	Active      bool    `db:"activo" json:"active"`
}
