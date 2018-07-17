package bootstrap

import (
	"github.com/josephniel/go-api/app/controllers"
	"github.com/josephniel/go-api/app/router"
	"github.com/josephniel/go-api/app/utils"
	"github.com/labstack/echo"
)

// RegisterRoutes encapsulates the route addition for the application
func RegisterRoutes(e *echo.Echo) {
	// initialize controllers first (i.e. so that the router can be prefilled)
	controllers.Init()
	// add routes to the application
	for _, route := range router.List() {
		utils.CallFuncByName(e, route.Method, route.Path, route.Controller)
	}
}
