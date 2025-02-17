package user

import (
	"database/sql"
	"gestor_de_usuario/internal/core/domain"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (userRepository *UserRepository) CreateUser(user *domain.User) (id interface{}, error error) {

	query := `
		INSERT INTO usuarios (
			usuario_id, nombre, apellido, email, contrasena, fecha_creacion, fecha_modificacion, activo
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := userRepository.Db.Exec(query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.Active,
	)
	if err != nil {
		return nil, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return id, nil

}

func (userRepository *UserRepository) GetUserByEmail(email string) (*domain.User, error) {

	query := `
		SELECT
			usuario_id, nombre, apellido, email, contrasena, fecha_creacion, fecha_modificacion, activo
		FROM usuarios
		WHERE email = ? and activo = 1
	`
	row := userRepository.Db.QueryRow(query, email)

	user := domain.User{}
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Active,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
