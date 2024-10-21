package http

import (
	"web-service-echo/app/http/middlewares"

	"github.com/labstack/echo/v4"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(middlewares.CorsDefault())
	e.Use(middlewares.BodyLimit())
	e.Use(middlewares.CsrfDefault())
	e.Use(middlewares.DefaultSecurity())
	//e.Use(middlewares.SetTimeout(30 * time.Second))
}
