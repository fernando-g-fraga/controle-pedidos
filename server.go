package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() *echo.Echo {

	e := echo.New()

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Listening :8000")

	return e
}
