package bootstrap

import (
	"fmt"

	"github.com/labstack/echo"
)

// Bootstrap is the container structure for the bootstrap functions
type Bootstrap struct {
	Server      *echo.Echo
	Environment *string
}

// Ignite calls in the bootstrap processes
func (b *Bootstrap) Ignite() {
	conf := GetConfiguration(b.Environment)
	SetupDB(conf)
	defer CloseDB()
	ApplyMiddlewares(b.Server, conf)
	RegisterRoutes(b.Server)
	b.Server.Logger.Fatal(b.Server.Start(fmt.Sprintf(":%d", conf.App.Port)))
}
