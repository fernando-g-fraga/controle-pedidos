package model

import "time"

/*
Table Produto
PK Código Int
Descrição String
Preço Float32

Table Departamento
PK Código Int
FK Código Produto INT
Descrição String

Join Table produtoPedido
quantidade integer
valorVenda float32

Table Pedido
PK numero int
data Date

*/

type produto struct {
	codigo    int
	descricao string
	preco     float32
}

type departamento struct {
	codigo         int
	codigo_produto int
	descricao      string
}

type produtoPedido struct {
	quantidade int
	valorVenda float32
}

type pedido struct {
	numero int
	data   time.Time
}
