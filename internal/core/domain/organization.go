package domain

import "time"

type Organization struct {
	ID          string    `db:"organizacion_id" json:"id"`
	Name        string    `db:"nombre" json:"name"`
	Description *string   `db:"descripcion" json:"description,omitempty"`
	CreatedAt   time.Time `db:"fecha_creacion" json:"created_at"`
	UpdatedAt   time.Time `db:"fecha_modificacion" json:"updated_at"`
	Active      bool      `db:"activo" json:"active"`
}
