package bootstrap

import (
	"fmt"

	"github.com/josephniel/go-api/app/model"
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

	db := DB{
		User:         conf.DB.User,
		Password:     conf.DB.Password,
		DatabaseName: conf.DB.Database,
		DatabaseObj:  &model.DB,
	}
	db.Start()
	defer db.Close()

	ApplyMiddlewares(b.Server, conf)
	RegisterRoutes(b.Server)
	b.Server.Logger.Fatal(b.Server.Start(fmt.Sprintf(":%d", conf.App.Port)))
}
