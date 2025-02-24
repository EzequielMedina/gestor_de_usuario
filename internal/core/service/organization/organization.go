package organization

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
	"gestor_de_usuario/internal/core/ports"
	"github.com/google/uuid"
	"time"
)

type OrganizationService struct {
	Repo ports.OrganizationRepository
}

func NewUserService(repo ports.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		Repo: repo,
	}
}

func builderNewOrganization(organization *request.OrganizationRequest) *domain.Organization {
	return &domain.Organization{
		ID:          uuid.New().String(),
		Name:        organization.Name,
		Description: organization.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Active:      true,
	}
}
