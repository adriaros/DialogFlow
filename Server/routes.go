package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// Root is the root route to test service health
func Root(c echo.Context) error {
	println("Root executed")
	return c.NoContent(http.StatusOK)
}

// Post is the root route to test service health
func Post(c echo.Context) error {
	println("Post executed")
	return c.NoContent(http.StatusOK)
}