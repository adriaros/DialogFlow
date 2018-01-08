package main

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/gommon/log"
)

// Root is the root route to test service health
func Root(c echo.Context) error {

	client, err := initDialogFlow()

	if err != nil {
		log.Print("Error initDialogFlow", err)
	}

	resp, err := dialogFlowQuery(client, "Hello")

	if err != nil {
		log.Print("Server Error: ", resp, " DialogFlow Error: ", err)
		return c.NoContent(http.StatusExpectationFailed)
	}

	println("Respuesta: ", resp)

	return c.NoContent(http.StatusOK)
}

// Post is the root route to test service health
func Post(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}