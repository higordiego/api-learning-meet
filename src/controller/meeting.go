package controller

import (
	"golang-meeting/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Created(c echo.Context) error {
	meeting := new(model.Meeting)
	if err := c.Bind(meeting); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := meeting.Created(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"created": true,
	})
}
