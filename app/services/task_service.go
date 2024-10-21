package services

import (
	task "web-service-echo/app/repositories/task"
	libraries "web-service-echo/libraries/db"

	"github.com/labstack/echo/v4"
)

func Init() {
	task.NewTaskRepository()
}

type TaskService interface {
	GetTaskData(ctx echo.Context) *libraries.Pagination
}

type TaskServiceImpl struct{}

func (s *TaskServiceImpl) GetTaskData(ctx echo.Context) *libraries.Pagination {
	repo := task.NewTaskRepository()

	return repo.GetAll(ctx)
}

func NewTaskService() TaskService {
	return &TaskServiceImpl{}
}
