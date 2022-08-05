package main

import (
	"github.com/labstack/echo/v4"
	controller "test.com/test/bobble/controller"
	database "test.com/test/bobble/data"
	service "test.com/test/bobble/service"
)

func main() {
	e := echo.New()
	database.GetConnection()
	database := database.NewBobbleDatabaseLayer()
	service := service.NewBobbleService(database)
	controller.NewBobbleController(e, service)
	e.Start(":8080")

}
