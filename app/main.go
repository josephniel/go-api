package main

import (
	"flag"
	"fmt"

	"github.com/josephniel/go-api/app/bootstrap"
	"github.com/labstack/echo"
)

func main() {
	env := flag.String("env", "test", "Application Environment")
	e := echo.New()

	conf := bootstrap.GetConfiguration(env)
	bootstrap.SetupDB(conf)
	defer bootstrap.CloseDB()
	bootstrap.ApplyMiddlewares(e, conf)
	bootstrap.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.App.Port)))
}
