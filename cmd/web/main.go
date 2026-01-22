package main

import (
	"greep-api-test/internal/config"
	"greep-api-test/internal/controller"
	"greep-api-test/internal/router"
	"greep-api-test/internal/server"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		createApp(),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func createApp() fx.Option {
	return fx.Options(
		config.Provide(),
		fx.Provide(zap.NewDevelopment),

		controller.Provide(),
		router.Provide(),
		server.Provide(),
	)
}
