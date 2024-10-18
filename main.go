package main

func main() {

	e := StartServer()

	//Calling Routes
	GetPedidosID(e)
	UpdatePedidos(e)
	DeletePedidos(e)
	GetPedidos(e)

}
