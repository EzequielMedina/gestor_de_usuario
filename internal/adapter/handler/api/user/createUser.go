package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/adapter/handler/api/response"
	"github.com/gin-gonic/gin"
)

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	var userRequest request.UserRequest
	if err := context.BindJSON(&userRequest); err != nil {
		response.ValidationError(context, err)
		return
	}

	newUser, err := userHandler.UserService.CreateUser(&userRequest)

	if err != nil {
		response.HandleError(context, err)
		return
	}

	resp := response.NewUserResponse(newUser)

	response.HandleSuccess(context, resp)

}
