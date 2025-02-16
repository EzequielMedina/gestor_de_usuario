package domain

type Permission struct {
	ID          string  `db:"permiso_id" json:"id"`
	Name        string  `db:"nombre" json:"name"`
	Description *string `db:"descripcion" json:"description,omitempty"`
	Active      bool    `db:"activo" json:"active"`
}
