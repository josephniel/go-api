package bootstrap

import (
	"github.com/josephniel/go-api/app/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ApplyMiddlewares encapsulates the all the middleware instatiation used in the application
func ApplyMiddlewares(e *echo.Echo, conf *config.Config) {
	applyBasicAuthMiddleware(e, conf)
	applyCORSMiddleware(e, conf)
}

func applyBasicAuthMiddleware(e *echo.Echo, conf *config.Config) {
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == conf.App.User && password == conf.App.Password {
			return true, nil
		}
		return false, nil
	}))
}

func applyCORSMiddleware(e *echo.Echo, conf *config.Config) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.CORS.AllowedOrigins,
		AllowMethods: conf.CORS.AllowedMethods,
	}))
}