package main

import (
	"github.com/josephniel/go-api/app/bootstrap"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	bootstrap.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
