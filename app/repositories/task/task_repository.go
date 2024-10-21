package task

import (
	"web-service-echo/app/models"
	"web-service-echo/config"
	libraries "web-service-echo/libraries/db"

	dbs "web-service-echo/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() TaskInterface {
	return &TaskRepository{
		db: dbs.GetDB(),
	}
}

func (t *TaskRepository) GetAll(e echo.Context) *libraries.Pagination {
	pagination := config.DefaultPagination(2, 1, "")
	tasks := make([]models.Task, 0)

	t.db.Scopes(libraries.Paginate(tasks, pagination, t.db)).Find(&tasks)
	pagination.Rows = tasks

	return pagination
}
