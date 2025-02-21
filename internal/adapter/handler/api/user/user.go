package user

import (
	"gestor_de_usuario/internal/core/ports"
)

type UserHandler struct {
	UserService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {

	return &UserHandler{
		UserService: userService,
	}

}
