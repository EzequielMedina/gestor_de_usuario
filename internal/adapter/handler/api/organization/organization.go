package organization

import "gestor_de_usuario/internal/core/ports"

type OrganizationHandler struct {
	OrganizationService ports.OrganizationService
}

func NewOrganizationHandler(organizationHandler ports.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		OrganizationService: organizationHandler,
	}
}
