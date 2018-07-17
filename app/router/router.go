package router

import "github.com/labstack/echo"

// Route is the DTO for a route object
type Route struct {
	Method     string
	Path       string
	Controller func(echo.Context) error
}

// Routes contain all the routes for the application
var Routes []Route

// Add is a generic function to add a route
func Add(method string, path string, controller func(echo.Context) error) {
	route := Route{
		Method:     method,
		Path:       path,
		Controller: controller,
	}
	Routes = append(Routes, route)
}
