package repository

import (
	"gestor_de_usuario/internal/core/domain"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	Db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{Db: db}
}

func (o *OrganizationRepository) CreateOrganization(organization *domain.Organization) error {
	result := o.Db.Table("organizaciones").Create(organization)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
