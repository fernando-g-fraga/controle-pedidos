package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fernandogfaga/controle-pedidos/model"
	"github.com/labstack/echo/v4"
)

func POSTpedidosHandler(c echo.Context) error {

	pedidoProduto := &model.PedidoProduto{}
	codigo_produto, _ := strconv.Atoi(c.FormValue("codigo_produto"))
	quantidade, _ := strconv.Atoi(c.FormValue("quantidade"))
	valorVenda, _ := strconv.ParseFloat(c.FormValue("valor_venda"), 64)
	pedidoProduto.Numero_pedido = model.CriaPedido()
	pedidoProduto.Codigo_produto = codigo_produto
	pedidoProduto.Quantidade = quantidade
	pedidoProduto.ValorVenda = valorVenda

	id_pedidoProduto := model.CriaPedidoProduto(*pedidoProduto)
	if id_pedidoProduto == 0 {
		log.Fatal("Error creating PedidoProduto")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	pedidoProduto.Numero = id_pedidoProduto
	return c.JSON(http.StatusOK, pedidoProduto)
}

func GETpedidosIDHandler(c echo.Context) error {
	id := c.FormValue("id")
	idconv, _ := strconv.Atoi(id)

	code, msg := model.GetPedidoID(idconv)

	if code == 400 {
		return c.JSON(code, map[string]string{"message": "Pedido não encontrado!"})
	}
	if code == 200 {
		return c.JSON(code, msg)
	}
	return c.JSON(500, map[string]string{"mensagem": "ocorreu um erro"})
}

func PUTpedidosHandler(c echo.Context) error {
	id := c.FormValue("id")
	data := c.FormValue("data")
	idconv, _ := strconv.Atoi(id)

	code, msg := model.AtualizaPedido(idconv, data)

	return c.JSON(code, msg)
}

func DELETEpedidosHandler(c echo.Context) error {
	id := c.FormValue("id")
	idconv, _ := strconv.Atoi(id)
	code, msg := model.DeletaPedido(idconv)

	return c.JSON(code, map[string]string{"message": msg})
}

func GETpedidosHandler(c echo.Context) error {
	return nil
}
