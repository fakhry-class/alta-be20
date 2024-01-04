package main

import (
	"fakhry/mvc/config"
	"fakhry/mvc/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
TODO:
Ubah tugas kemarin ke MVC structure

tambahkan endpoint
GET /users/:user_id/products --> untuk mendapatkan list product yang dimiliki oleh user dengan id tertentu.
/users/9/products
*/
func main() {
	fmt.Println("running")
	config.InitDB()
	config.InitialMigration()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	routes.InitRoute(e)

	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
