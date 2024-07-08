package cmd

import (
	"context"
	"os"
	"time"
	"ups02/cmd/server"
	"ups02/config"
	"ups02/provider"

	"github.com/sirupsen/logrus"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()
	config.LoadEnvirontment()

	db, err := config.SetDatabase()
	if err != nil {
		logrus.Fatalf("Error setting up dastabase %v", err)
	}

	provider.NewProvider(db, server)
	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); err != nil {
			logrus.Errorf("Error Starting the server %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("Server stopped")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Error the stopping server %v", err)
	}

}
