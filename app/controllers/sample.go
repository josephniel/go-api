package controllers

import (
	"net/http"

	"github.com/josephniel/go-api/app/operations"
	"github.com/josephniel/go-api/app/router"
	"github.com/josephniel/go-api/app/schema"
	"github.com/labstack/echo"
)

func getUser(context echo.Context) error {
	sampleStruct := new(schema.Sample)
	if err := context.Bind(sampleStruct); err != nil {
		context.JSON(http.StatusBadRequest, schema.Error{
			Message: err.Error(),
		})
		return err
	}
	if err := context.Validate(sampleStruct); err != nil {
		context.JSON(http.StatusUnprocessableEntity, schema.Error{
			Message: err.Error(),
		})
		return err
	}
	response, err := operations.GetUser(sampleStruct)
	if err != nil {
		context.JSON(http.StatusInternalServerError, schema.Error{
			Message: err.Error(),
		})
		return err
	}
	return context.JSON(http.StatusOK, response)
}

func init() {
	router.Post("/users", getUser)
}
