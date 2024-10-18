package main

import (
	"log"

	"github.com/fernandogfaga/controle-pedidos/model"
)

func main() {

	e := StartServer()
	log.Println("Server started!")
	db := model.StartDB()

	//Calling Routes
	GetPedidosID(e)
	UpdatePedidos(e)
	DeletePedidos(e)
	GetPedidos(e)
	log.Println("Routes Ready to roll")

	//Executing SQL First Commands
	row := model.CreateTables(db)
	log.Println(row)

}
