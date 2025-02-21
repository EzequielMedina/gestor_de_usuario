package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/response"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetUserByEmail(c *gin.Context) {

	//obtenemos de la url el parametro email
	email := c.Query("email")

	user, err := h.UserService.GetUserByEmail(email)

	if err != nil {
		response.HandleError(c, err)
		return
	}

	resp := response.NewUserResponse(user)
	response.HandleSuccess(c, resp)

}
