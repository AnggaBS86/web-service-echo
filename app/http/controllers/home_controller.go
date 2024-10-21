package controllers

import (
	"web-service-echo/app/services"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
	*Controller
}

var service = services.NewTaskService()

func (c *HomeController) Hello(ctx echo.Context) error {
	return c.Success(ctx, "Hello, World!")
}

func (c *HomeController) GetTasks(ctx echo.Context) error {
	tasks := service.GetTaskData(ctx)
	return c.Success(ctx, tasks)
}
