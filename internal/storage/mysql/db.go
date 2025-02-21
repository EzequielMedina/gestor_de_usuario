package mysql

import (
	"fmt"
	"gestor_de_usuario/internal/adapter/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config *config.DB) (*gorm.DB, error) {
	// Open the connection
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
