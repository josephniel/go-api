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
		Port:        flag.Int64("port", 8080, "Application Port"),
	}
	flag.Parse()
	server.Ignite()
}
