package routes

import (
	"github.com/kzw200015/picbed/assets"
	"github.com/kzw200015/picbed/config"
	"github.com/kzw200015/picbed/handlers"
	"github.com/kzw200015/picbed/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echo.New()

func Get() *echo.Echo {
	return e
}

func init() {
	e.HTTPErrorHandler = handlers.HandleError

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] INFO ${latency_human} ${status} ${method} ${uri}" + "\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CORS())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: assets.GetFS(),
	}))

	images := e.Group(config.IMG_ROUTE_PREFIX)
	{
		images.Use(middlewares.Referer)
		images.Use(middleware.Static(config.Get().ImgSrc))
	}

	api := e.Group("/api")
	{
		api.POST("/upload", handlers.HandleUpload)
		api.DELETE("/remove/:id", handlers.HandleRemove)
	}
}
