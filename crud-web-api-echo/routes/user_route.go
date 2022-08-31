package routes

import (
	"echo-mongo-api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", (&controllers.Handler{}).CreateStudent)
	e.GET("/:userId", (&controllers.Handler{}).GetAStudent)
	e.PUT("/:userId", (&controllers.Handler{}).EditAStudent)
	e.DELETE("/:userId", (&controllers.Handler{}).DeleteAStudent)
	e.GET("/all", (&controllers.Handler{}).GetAllStudents)
}
