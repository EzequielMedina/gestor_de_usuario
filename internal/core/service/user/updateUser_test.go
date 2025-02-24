package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"testing"
)

type updateTestedInput struct {
	user *request.UserUpdateRequest
}
type updateTestedOutput struct {
	err error
}

func TestUserService_UpdateUser(t *testing.T) {

}
