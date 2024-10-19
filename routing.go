package main

import (
	"github.com/fernandogfaga/controle-pedidos/controller"
	"github.com/labstack/echo/v4"
)

// Handlers

func PostPedidos(e *echo.Echo) {
	e.POST("/add_pedidos", controller.POSTpedidosHandler)

}

// Routes
func GetPedidosID(e *echo.Echo) {
	e.GET("/pedidos", controller.GETpedidosIDHandler)

}

func UpdatePedidos(e *echo.Echo) {
	e.PUT("/update_pedidos", controller.PUTpedidosHandler)

}

func DeletePedidos(e *echo.Echo) {
	e.DELETE("/delete_pedidos", controller.DELETEpedidosHandler)

}

func GetPedidos(e *echo.Echo) {
	e.POST("/get_pedidos", controller.GETpedidosHandler)

}
