package router

import (
	"greep-api-test/internal/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func newRouter(c *controller.Controller) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	router.Static("/static", "./ui/static")
	router.LoadHTMLFiles("ui/html/home.tmpl")

	router.GET("/home", c.Home)
	router.POST("/api/v1/get-settings", c.GetSettings)
	router.POST("/api/v1/get-state-instance", c.GetStateInstance)
	router.POST("/api/v1/send-message", c.SendMessage)
	router.POST("/api/v1/send-file-by-url", c.SendFileByUrl)

	return router, nil
}

func Provide() fx.Option {
	return fx.Provide(newRouter)
}
