package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
	"gestor_de_usuario/internal/core/ports"
)

type UserService struct {
	Repo  ports.UserRepository
	Utils ports.UtilService
}

func NewUserService(repo ports.UserRepository, util ports.UtilService) *UserService {
	return &UserService{
		Repo:  repo,
		Utils: util,
	}
}

func ValidateCreateUser(user *request.UserRequest, s *UserService) (bool, error) {
	//validar que el email no sea vacio

	if user.Email == "" {
		return true, domain.ErrEmailRequired
	}

	//validar que el email sea un email valido

	if !s.Utils.IsValidEmail(user.Email) {
		return true, domain.ErrInvalidEmail
	}

	//validar que el email no exista en la base de datos

	userDbMail, err := s.Repo.GetUserByEmail(user.Email)

	if err != nil {
		return true, domain.ErrInternal
	}

	if userDbMail != nil {
		return true, domain.ErrEmailAlreadyExists
	}

	//validar que el nombre no sea vacio

	if user.Name == "" {
		return true, domain.ErrNameRequired
	}

	//validar que el apellido no sea vacio

	if user.LastName == "" {
		return true, domain.ErrLastNameRequired
	}

	//validar que la contraseña no sea vacia

	if user.Password == "" {
		return true, domain.ErrPasswordRequired
	}

	//validamos que la contraseña sea segura mas de 8 digitos alfanumerica con un signo

	if !s.Utils.IsValidPassword(user.Password) {
		return true, domain.ErrInvalidPassword
	}
	return false, nil
}
