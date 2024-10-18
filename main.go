package main

import (
	"log"
	"net/http"

	"github.com/fernandogfaga/controle-pedidos/model"
)

func main() {

	//Executing SQL Model
	model.Config()

	//Runing server
	e := StartServer()

	//Calling Routes
	PostPedidos(e)
	GetPedidosID(e)
	UpdatePedidos(e)
	DeletePedidos(e)
	GetPedidos(e)
	log.Println("Routes Ready to roll")

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Listening :8000")
}
