package user

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
)

func (s *UserService) UpdateUser(user *request.UserUpdateRequest) error {
	userDb, err := s.Repo.GetById(user.ID)

	if err != nil {
		return err
	}

	if err := UpdateEmail(userDb, user.Email, s); err != nil {
		return err
	}

	UpdateNames(userDb, user.Name, user.LastName)

	if err := UpdatePassword(userDb, user.Password, s); err != nil {
		return err
	}

	err = s.Repo.UpdateUser(userDb)
	if err != nil {
		return err
	}

	return nil
}
