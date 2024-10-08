package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsDefault() echo.MiddlewareFunc {
	return middleware.CORS()
}
