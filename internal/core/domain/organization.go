package domain

import "time"

type Organization struct {
	ID          string    `gorm:"column:organizacion_id;primaryKey" json:"id"`
	Name        string    `gorm:"column:nombre" json:"name"`
	Description string    `gorm:"column:descripcion" json:"description"`
	CreatedAt   time.Time `gorm:"column:fecha_creacion" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:fecha_modificacion" json:"updated_at"`
	Active      bool      `gorm:"column:activo" json:"active"`
}
