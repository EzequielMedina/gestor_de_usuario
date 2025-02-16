package response

import (
	"gestor_de_usuario/internal/core/domain"
	"time"
)

// userResponse represents a user response body
type userResponse struct {
	ID        string    `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"test@example.com"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

// newUserResponse is a helper function to create a response body for handling user data
func NewUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.FirstName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
