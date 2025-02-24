package organization

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
)

func (s *OrganizationService) CreateOrganization(organization *request.OrganizationRequest) (*domain.Organization, error) {
	if organization.Name == "" {
		return nil, domain.ErrNameRequired
	}
	newOrganization := builderNewOrganization(organization)

	err := s.Repo.CreateOrganization(newOrganization)

	if err != nil {
		return nil, domain.ErrCreateOrganization
	}

	return newOrganization, nil
}
