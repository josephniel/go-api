package bootstrap

import (
	_ "github.com/josephniel/go-api/app/controller" // import controller so all contents are initialized first
	"github.com/josephniel/go-api/app/router"
	"github.com/josephniel/go-api/app/util"
	"github.com/labstack/echo"
)

// RegisterRoutes encapsulates the route addition for the application
func RegisterRoutes(e *echo.Echo) {
	for _, route := range router.List() {
		util.CallFuncByName(e, route.Method, route.Path, route.Controller)
	}
}
