package user

import (
	"database/sql"
	"gestor_de_usuario/internal/core/domain"
	"time"
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
	var createdAt, updatedAt []byte

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&createdAt,
		&updatedAt,
		&user.Active,
	)
	if err != nil {
		return nil, err
	}

	// Convertir los valores de []byte a time.Time
	user.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
	user.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
