package main

import (
	"github.com/labstack/echo/v4"
)

func StartServer() *echo.Echo {
	e := echo.New()
	return e
}
