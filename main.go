package main

import (
	"log"
	"net/http"

	"github.com/fernandogfaga/controle-pedidos/model"
)

func main() {

	//Executing SQL First Commands
	db := model.StartDB()
	row := model.CreateTables(db)
	log.Println(row)

	//creating mocks (1 time only)
	model.InsertMockProdutos(db)

	//Runing server
	e := StartServer()

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Listening :8000")

	//Calling Routes
	GetPedidosID(e)
	UpdatePedidos(e)
	DeletePedidos(e)
	GetPedidos(e)
	log.Println("Routes Ready to roll")

}
