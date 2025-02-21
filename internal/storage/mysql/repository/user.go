package user

import (
	"gestor_de_usuario/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (userRepository *UserRepository) CreateUser(user *domain.User) (id interface{}, error error) {

	result := userRepository.Db.Table("Usuarios").Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user.ID, nil

}

func (userRepository *UserRepository) GetUserByEmail(email string) (*domain.User, error) {

	var user domain.User
	result := userRepository.Db.Table("Usuarios").Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return nil, domain.ErrDataNotFound
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrEmailNotFound
	}

	return &user, nil
}
