package routes

import (
	"web-service-echo/app/http/controllers"

	"github.com/labstack/echo/v4"
)

var GlobalRoutes = func(e *echo.Echo) {
	//Routes group "api/v1"
	routerGroup := e.Group("/api/v1")

	homeController := new(controllers.HomeController)
	//hello world
	routerGroup.GET("/hello", homeController.Hello)

	//list tasks
	taskRoute := routerGroup.Group("/tasks")
	taskRoute.GET("/", homeController.GetTasks)
}
