package routers

import (
	"fakhry/clean-arch/app/middlewares"
	"fakhry/clean-arch/features/user/data"
	_userHandler "fakhry/clean-arch/features/user/handler"
	_userService "fakhry/clean-arch/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	userData := data.New(db)
	// userData := _userData.NewRaw(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)
	// define routes/ endpoint
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUsers, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.Update)

}
