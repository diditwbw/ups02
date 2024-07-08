package routes

import (
	"ups02/cmd/server"
	"ups02/internals/handler"
	"ups02/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterUserRoutes(server server.Ginserver, userHandler *handler.Userhandler) {
	server.RegisterGroupRoute("api/v1/user", []interfacesx.RouteDefinition{
		{Method: "POST", Path: "/register", Handler: userHandler.CreateUser},
	}, func(c *gin.Context) {
		logrus.Infof("Request on %s", c.Request.URL.Path)
	})
}
