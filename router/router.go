package router

import (
	"github.com/labstack/echo"
)

// Route is the schema used for the application routes DTO
type Route struct {
	Method     string
	Path       string
	Controller func(echo.Context) error
}

var routes []Route

// Add includes a route to the route list
func Add(method string, path string, controller func(echo.Context) error) {
	route := Route{
		Method:     method,
		Path:       path,
		Controller: controller,
	}
	routes = append(routes, route)
}

// List returns all the routes of the application
func List() []Route {
	return routes
}
