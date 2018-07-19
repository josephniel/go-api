package main

import (
	"flag"

	"github.com/josephniel/go-api/app/bootstrap"
	"github.com/labstack/echo"
)

func main() {
	server := bootstrap.Bootstrap{
		Environment: flag.String("env", "test", "Application Environment"),
		Server:      echo.New(),
	}
	server.Ignite()
}
