package user

import (
	"gestor_de_usuario/internal/adapter/handler/api"
	"gestor_de_usuario/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {

	return &UserHandler{
		UserService: userService,
	}

}

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	var userRequest UserRequest
	if err := context.BindJSON(&userRequest); err != nil {
		api.ValidationError(context, err)
		return
	}

	newUser, err := userHandler.UserService.CreateUser(&userRequest)

	if err != nil {
		api.HandleError(context, err)
		return
	}

	response := newUserResponse(newUser)

	api.HandleSuccess(context, response)

	api.HandleSuccess(context, newUser)

}
