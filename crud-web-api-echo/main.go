package main

import (
	"echo-mongo-api/configs"
	"echo-mongo-api/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()
	fmt.Println("connected to server")
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":8000"))
}
