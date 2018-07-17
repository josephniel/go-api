package main

import (
	"github.com/josephniel/api/controllers"
	"github.com/josephniel/api/router"
	"github.com/josephniel/api/utils"
	"github.com/labstack/echo"
)

func main() {
	controllers.Init()

	e := echo.New()
	for _, route := range router.List() {
		utils.CallFuncByName(e, route.Method, route.Path, route.Controller)
	}
	e.Logger.Fatal(e.Start(":1323"))
}
