package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Ginserver interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
}

type GinServerBuilder struct {
}

type ginServer struct {
	engine *gin.Engine
	server *http.Server
}

func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

func (b *GinServerBuilder) Build() Ginserver {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}

	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Listening %s \n", err)
		}
	}()

	logrus.Infof("https server is running on port %s", httpAddress)
	return nil
}

func (gs *ginServer) Shutdown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown %s", err)
	}

	logrus.Info("Server is Exiting")

	return nil
}
