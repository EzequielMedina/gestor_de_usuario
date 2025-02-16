package api

import (
	configuration "gestor_de_usuario/internal/adapter/config"
	"gestor_de_usuario/internal/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	config, err := configuration.New()
	if err != nil {
		log.Fatalf("Could not load the configuration: %v", err)
	}

	// Connect to the database
	db, err := mysql.Connect(config.DB)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()
}
