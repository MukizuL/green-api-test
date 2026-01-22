package server

import (
	"context"
	"errors"
	"greep-api-test/internal/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func newHTTPServer(lc fx.Lifecycle, cfg *config.Config, router *gin.Engine, logger *zap.Logger) *http.Server {
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router.Handler(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server", zap.String("addr", cfg.Addr))

			var err error

			go func() {
				err = srv.ListenAndServe()
			}()

			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func Provide() fx.Option {
	return fx.Provide(newHTTPServer)
}
