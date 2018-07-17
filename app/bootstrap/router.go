package bootstrap

import (
	"github.com/josephniel/go-api/app/controllers"
	"github.com/josephniel/go-api/app/router"
	"github.com/josephniel/go-api/app/utils"
	"github.com/labstack/echo"
)

// RegisterRoutes encapsulates the route addition for the application
func RegisterRoutes(e *echo.Echo) {
	controllers.Init()
	for _, route := range router.Routes {
		utils.CallFuncByName(e, route.Method, route.Path, route.Controller)
	}
}
