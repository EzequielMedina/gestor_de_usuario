package domain

import "time"

type User struct {
	ID        string    `db:"usuario_id" json:"id"`
	FirstName string    `db:"nombre" json:"first_name"`
	LastName  string    `db:"apellido" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"contrasena" json:"password"`
	CreatedAt time.Time `db:"fecha_creacion" json:"created_at"`
	UpdatedAt time.Time `db:"fecha_modificacion" json:"updated_at"`
	Active    bool      `db:"activo" json:"active"`
}
