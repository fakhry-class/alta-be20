package routes

import (
	"fakhry/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	// create a new echo instance

	// define routes/ endpoint
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users", controllers.GetAllUserController)
	e.PUT("/users/:user_id", controllers.UpdateUserByIdController)

	// return e
}
