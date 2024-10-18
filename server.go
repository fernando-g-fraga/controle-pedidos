package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() *echo.Echo {

	e := echo.New()

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	return e
}
