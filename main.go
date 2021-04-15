package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"websocket-example/controller"
)

func main() {
	go controller.H.Run()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws/:roomId", controller.ServeWs)
	e.Logger.Fatal(e.Start(":1323"))
}
