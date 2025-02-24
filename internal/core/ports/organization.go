package ports

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
)

type OrganizationRepository interface {
	CreateOrganization(organization *domain.Organization) error
}

type OrganizationService interface {
	CreateOrganization(organization *request.OrganizationRequest) (*domain.Organization, error)
}
