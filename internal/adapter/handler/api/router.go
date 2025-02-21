package api

import (
	"gestor_de_usuario/internal/adapter/config"
	"gestor_de_usuario/internal/adapter/handler/api/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	userHandler user.UserHandler,
) (*Router, error) {
	// Disable debug mode in production
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())

	//se agregan los endpoints

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/create", userHandler.CreateUser)
			users.GET("getByUserEmail", userHandler.GetUserByEmail)
		}

	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
