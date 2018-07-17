package controllers

import (
	"net/http"
	"strconv"

	"github.com/josephniel/go-api/app/router"
	"github.com/josephniel/go-api/app/schema"
	"github.com/labstack/echo"
)

type sample struct {
	ID int `json:"id"`
}

func getUser(context echo.Context) error {
	id, err := strconv.ParseInt(context.Param("id"), 10, 32)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, schema.Error{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, sample{
		ID: int(id),
	})
}

func init() {
	router.Get("/users/:id", getUser)
}
