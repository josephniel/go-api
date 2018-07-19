package bootstrap

import (
	"github.com/go-playground/validator"
	"github.com/josephniel/go-api/app/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ApplyMiddlewares encapsulates the all the middleware instatiation used in the application
func ApplyMiddlewares(e *echo.Echo, conf *config.Config) {
	applyBasicAuthMiddleware(e, conf)
	applyRequestValidator(e)
	applyLogger(e)
}

func applyBasicAuthMiddleware(e *echo.Echo, conf *config.Config) {
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == conf.Auth.User && password == conf.Auth.Password {
			return true, nil
		}
		return false, nil
	}))
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func applyRequestValidator(e *echo.Echo) {
	e.Validator = &customValidator{validator: validator.New()}
}

func applyLogger(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
}
