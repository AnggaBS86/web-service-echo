package task

import (
	"web-service-echo/app/models"
	"web-service-echo/config"
	"web-service-echo/db"
	libraries "web-service-echo/libraries/db"

	"github.com/labstack/echo/v4"
)

type TaskRepository struct{}

func NewTaskRepository() TaskInterface {
	return &TaskRepository{}
}

func (t *TaskRepository) GetAll(e echo.Context) *libraries.Pagination {
	pagination := config.DefaultPagination(2, 1, "")
	tasks := make([]models.Task, 0)
	db := db.GetDB()
	db.Scopes(libraries.Paginate(tasks, pagination, db)).Find(&tasks)
	pagination.Rows = tasks

	return pagination
}
