package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// Root is the root route to test service health
func Root(c echo.Context) error {

	initDialogFlow()

	return c.NoContent(http.StatusOK)
}

// Post is the root route to test service health
func Post(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}