package task

import (
	libraries "web-service-echo/libraries/db"

	"github.com/labstack/echo/v4"
)

type TaskInterface interface {
	GetAll(e echo.Context) *libraries.Pagination
}
