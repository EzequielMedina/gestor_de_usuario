package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/adapter/handler/api/response"

	"github.com/gin-gonic/gin"
)

func (userHandler *UserHandler) UpdateUser(context *gin.Context) {
	var userUpdateRequest request.UserUpdateRequest
	if err := context.BindJSON(&userUpdateRequest); err != nil {
		response.ValidationError(context, err)
		return
	}

	userUpdateRequest.ID = context.Param("id")

	err := userHandler.UserService.UpdateUser(&userUpdateRequest)

	if err != nil {
		response.HandleError(context, err)
		return
	}

	response.HandleSuccess(context, nil)

}
