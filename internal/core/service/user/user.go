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
	//validar que no existe

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
