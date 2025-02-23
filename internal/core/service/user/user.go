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
	userDbMail, err := ValidarEmail(user.Email, s)

	if err != nil {
		return true, err
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

func ValidarEmail(email string, s *UserService) (*domain.User, error) {
	//validar que el email no sea vacio

	if email == "" {
		return nil, domain.ErrEmailRequired
	}

	//validar que el email sea un email valido

	if !s.Utils.IsValidEmail(email) {
		return nil, domain.ErrInvalidEmail
	}

	//validar que el email no exista en la base de datos

	userDbMail, err := s.Repo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if userDbMail != nil {
		return userDbMail, nil
	}

	return nil, nil
}
func UpdateEmail(userDb *domain.User, email string, s *UserService) error {
	if email == "" {
		return nil
	}

	userEmailDB, err := s.Repo.GetUserByEmail(email)

	if err != nil {
		return err
	}

	if userEmailDB != nil {
		return domain.ErrEmailAlreadyExists
	}

	userDb.Email = email
	return nil
}

func UpdateNames(userDb *domain.User, name string, lastName string) {
	if name != "" {
		userDb.FirstName = name
	}
	if lastName != "" {
		userDb.LastName = lastName
	}
}

func UpdatePassword(userDb *domain.User, password string, s *UserService) error {
	if password == "" {
		return nil
	}

	if !s.Utils.IsValidPassword(password) {
		return domain.ErrInvalidPassword
	}

	hashPassword, err := s.Utils.HashPassword(password)

	if err != nil {
		return err
	}

	userDb.Password = hashPassword
	return nil
}
