package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
	"gestor_de_usuario/internal/core/ports"
	"gestor_de_usuario/internal/core/util"
	"github.com/google/uuid"
	"time"
)

type UserService struct {
	Repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(user *request.UserRequest) (*domain.User, error) {

	done, errs := validateCreateUser(user, s)
	if done {
		return nil, errs
	}

	// set date
	var newUser domain.User
	newUser.ID = uuid.New().String()
	hashPassword, err := util.HashPassword(user.Password)

	if err != nil {
		return nil, domain.ErrInternal
	}

	newUser.CreatedAt = time.Now().UTC()
	newUser.Active = true
	newUser.Password = hashPassword
	newUser.Email = user.Email
	newUser.FirstName = user.Name
	newUser.LastName = user.LastName
	newUser.UpdatedAt = time.Now().UTC()
	_, err = s.Repo.CreateUser(&newUser)

	if err != nil {
		return nil, domain.ErrInternal
	}

	return &newUser, nil

}

func validateCreateUser(user *request.UserRequest, s *UserService) (bool, error) {
	//validar que el email no sea vacio

	if user.Email == "" {
		return true, domain.ErrEmailRequired
	}

	//validar que el email sea un email valido

	if !util.IsValidEmail(user.Email) {
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

	if !util.IsValidPassword(user.Password) {
		return true, domain.ErrInvalidPassword
	}
	return false, nil
}
