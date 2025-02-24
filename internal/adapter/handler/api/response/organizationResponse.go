package response

import (
	"gestor_de_usuario/internal/core/domain"
	"time"
)

type organizationResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	createAt time.Time `json:"created_at"`
	updateAt time.Time `json:"updated_at"`
}

func NewOrganizationResponse(organization *domain.Organization) organizationResponse {
	return organizationResponse{
		ID:       organization.ID,
		Name:     organization.Name,
		createAt: organization.CreatedAt,
		updateAt: organization.UpdatedAt,
	}
}
