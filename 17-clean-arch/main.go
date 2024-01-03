package main

import (
	"fakhry/clean-arch/app/configs"
	"fakhry/clean-arch/app/databases"
	"fakhry/clean-arch/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbSql := databases.InitDBMysql(cfg)

	e := echo.New()
	routers.InitRouter(dbSql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
