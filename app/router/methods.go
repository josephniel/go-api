package router

import "github.com/labstack/echo"

// Get adds a GET method given a path and a controller
func Get(path string, controller func(echo.Context) error) {
	Add(echo.GET, path, controller)
}

// Post adds a POST method given a path and a controller
func Post(path string, controller func(echo.Context) error) {
	Add(echo.POST, path, controller)
}

// Patch adds a PATCH method given a path and a controller
func Patch(path string, controller func(echo.Context) error) {
	Add(echo.PATCH, path, controller)
}

// Delete adds a DELETE method given a path and a controller
func Delete(path string, controller func(echo.Context) error) {
	Add(echo.DELETE, path, controller)
}
