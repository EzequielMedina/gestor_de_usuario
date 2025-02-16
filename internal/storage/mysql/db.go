package mysql

import (
	"database/sql"
	"fmt"
	"gestor_de_usuario/internal/adapter/config"
)

func Connect(config *config.DB) (*sql.DB, error) {
	// Open the connection
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open(config.Connection, dns)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
