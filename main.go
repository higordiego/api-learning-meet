package main

import (
	"golang-meeting/src/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/meeting", controller.Created)
	e.Logger.Fatal(e.Start(":3000"))
}
