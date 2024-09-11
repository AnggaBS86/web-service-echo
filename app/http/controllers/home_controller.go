package controllers

import (
	"web-service-echo/app/repositories/task"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
	*Controller
}

var service = task.NewTaskRepository()

func (c *HomeController) Hello(ctx echo.Context) error {
	return c.Success(ctx, "Hello, World!")
}

func (c *HomeController) GetTasks(ctx echo.Context) error {
	tasks := service.GetAll(ctx)
	return c.Success(ctx, tasks)
}
