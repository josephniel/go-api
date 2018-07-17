package router

import "github.com/labstack/echo"

// Route is the DTO for a route object
type Route struct {
	Method     string
	Path       string
	Controller func(echo.Context) error
}

var routes []Route

// Add is a generic function to add a route
func Add(method string, path string, controller func(echo.Context) error) {
	route := Route{
		Method:     method,
		Path:       path,
		Controller: controller,
	}
	routes = append(routes, route)
}

// List returns all the routes for the application
func List() []Route {
	return routes
}
