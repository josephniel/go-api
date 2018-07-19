package controllers

import (
	"net/http"
	"strconv"

	"github.com/josephniel/go-api/app/models"
	"github.com/josephniel/go-api/app/operations"
	"github.com/josephniel/go-api/app/router"
	"github.com/josephniel/go-api/app/schema"
	"github.com/labstack/echo"
)

func getUser(context echo.Context) error {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return context.JSON(http.StatusBadRequest, schema.Message{
			Message: err.Error(),
		})
	}
	response, err := operations.GetUser(id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, schema.Message{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, response)
}

func addUser(context echo.Context) error {
	user := new(models.User)
	if err := context.Bind(user); err != nil {
		return context.JSON(http.StatusBadRequest, schema.Message{
			Message: err.Error(),
		})
	}
	if err := context.Validate(user); err != nil {
		return context.JSON(http.StatusUnprocessableEntity, schema.Message{
			Message: err.Error(),
		})
	}
	err := operations.AddUser(user)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, schema.Message{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, schema.Message{
		Message: "User successfully added!",
	})
}

func init() {
	router.Get("/users/:id", getUser)
	router.Post("/users", addUser)
}
