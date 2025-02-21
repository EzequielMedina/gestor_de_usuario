package user

import "gestor_de_usuario/internal/core/domain"

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {

	userDbMail, err := ValidarEmail(email, s)

	if err != nil {
		return nil, err
	}
	
	return userDbMail, nil
}
