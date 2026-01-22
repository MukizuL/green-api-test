package controller

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger
}

func newController(logger *zap.Logger) *Controller {
	return &Controller{
		logger: logger,
	}
}

func Provide() fx.Option {
	return fx.Provide(newController)
}
