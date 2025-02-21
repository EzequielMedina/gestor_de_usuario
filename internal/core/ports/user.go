package ports

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (id interface{}, error error)
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService interface {
	CreateUser(user *request.UserRequest) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
