package domain

import "time"

type User struct {
	ID        string    `gorm:"column:usuario_id;primaryKey" json:"id"`
	FirstName string    `gorm:"column:nombre" json:"first_name"`
	LastName  string    `gorm:"column:apellido" json:"last_name"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Password  string    `gorm:"column:contrasena" json:"password"`
	CreatedAt time.Time `gorm:"column:fecha_creacion" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:fecha_modificacion" json:"updated_at"`
	Active    bool      `gorm:"column:activo" json:"active"`
}
