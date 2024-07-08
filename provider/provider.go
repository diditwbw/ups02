package provider

import (
	"ups02/cmd/server"
	"ups02/internals/handler"
	"ups02/internals/repository"
	"ups02/internals/routes"
	"ups02/internals/services"

	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.Ginserver) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userhandler := handler.NewUserhandler(userService)
	routes.RegisterUserRoutes(server, userhandler)
}
