package main

import (
	"fmt"
	"log"
	"os"
	httpmiddleware "web-service-echo/app/http"
	"web-service-echo/db"
	"web-service-echo/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "gorm.io/driver/mysql"
)

func main() {
	//get conf from .env file
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println(err.Error())
	}

	//database initialization
	db.Init()

	//initialize *echo.Echo
	e := echo.New()

	//initilize middlewares that used by this app
	httpmiddleware.InitMiddlewares(e)

	//initialize routes
	routes.GlobalRoutes(e)

	//get port
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	//run the server
	e.Logger.Fatal(e.Start(port))
}
