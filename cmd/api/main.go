package main

import (
	configuration "gestor_de_usuario/internal/adapter/config"
	"gestor_de_usuario/internal/adapter/handler/api"
	userHandler "gestor_de_usuario/internal/adapter/handler/api/user"
	userService "gestor_de_usuario/internal/core/service/user"
	"gestor_de_usuario/internal/core/util"
	"gestor_de_usuario/internal/storage/mysql"
	userResitory "gestor_de_usuario/internal/storage/mysql/repository"
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

	//Dependency Injection
	utilSrv := util.NewUtilService()

	userRepor := userResitory.NewUserRepository(db)
	userSrv := userService.NewUserService(userRepor, utilSrv)
	userHand := userHandler.NewUserHandler(userSrv)

	router, err := api.NewRouter(config.Http, *userHand)
	if err != nil {
		log.Fatalf("Could not create the router: %v", err)
	}

	listenAddr := config.Http.URL + ":" + config.Http.Port

	// Start the server
	err = router.Serve(listenAddr)

	if err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}

}
