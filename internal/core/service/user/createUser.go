package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
	"github.com/google/uuid"
	"time"
)

func (s *UserService) CreateUser(user *request.UserRequest) (*domain.User, error) {

	done, errs := ValidateCreateUser(user, s)
	if done {
		return nil, errs
	}

	// set date
	var newUser domain.User
	newUser.ID = uuid.New().String()
	hashPassword, err := s.Utils.HashPassword(user.Password)

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
