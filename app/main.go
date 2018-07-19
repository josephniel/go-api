package main

import (
	"flag"
	"fmt"

	"github.com/go-pg/pg"
	"github.com/josephniel/go-api/app/bootstrap"
	"github.com/josephniel/go-api/app/models"
	"github.com/labstack/echo"
)

func main() {
	env := flag.String("env", "test", "Application Environment")
	e := echo.New()

	conf := bootstrap.GetConfiguration(env)

	db := pg.Connect(&pg.Options{
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
	})
	models.DB = db
	defer db.Close()

	bootstrap.ApplyMiddlewares(e, conf)
	bootstrap.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.App.Port)))
}
