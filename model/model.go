package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

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

type Produto struct {
	Codigo    int     `json:"codigo"`
	Descricao string  `json:"descricao"`
	Preco     float32 `json:"preco"`
	Nome      string  `json:"nome"`
}

type Departamento struct {
	codigo         int
	codigo_produto int
	descricao      string
}

type PedidoProduto struct {
	Numero         int64
	Codigo_produto int
	Numero_pedido  int64
	Quantidade     int
	ValorVenda     float64
}

type Pedido struct {
	numero int
	data   time.Time
}

var db = StartDB()

func Config() {
	CreateTables()
	CreateMock()
}

func CreateTables() sql.Result {

	query := `
	CREATE TABLE IF NOT EXISTS pedido(
	 	numero SERIAL NOT NULL PRIMARY KEY,
	 	data DATE
	 );
	
	CREATE TABLE IF NOT EXISTS produto(
		codigo SERIAL NOT NULL PRIMARY KEY,
		descricao VARCHAR (255),
		preco NUMERIC,
		NOME VARCHAR(70)
	);

	CREATE TABLE IF NOT EXISTS  departamento(
		codigo SERIAL NOT NULL	PRIMARY KEY,
		codigo_produto INT NOT NULL,
		descricao VARCHAR (255),
		FOREIGN KEY (codigo_produto) REFERENCES Produto(codigo)
	);

	CREATE TABLE IF NOT EXISTS pedidoProduto(
		codigo SERIAL NOT NULL PRIMARY KEY,
		codigo_produto INT NOT NULL,
		numero_pedido INT NOT NULL,
		quantidade INT,
		valorVenda NUMERIC,
		FOREIGN KEY (codigo_produto) REFERENCES Produto(codigo),
		FOREIGN KEY (numero_pedido) REFERENCES Pedido(numero)
	);
	`
	row, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating the tables", err)
	}

	return row
}

func InsertMockProdutos() {

	produtos := CreateMock()

	for _, valor := range produtos {

		_, err := db.Exec("INSERT INTO produto VALUES($1,$2,$3,$4);", valor.Codigo, valor.Descricao, valor.Preco, valor.Nome)
		if err != nil {
			log.Fatal("Error creating the mock", err)
		}

	}

}

func CriaPedido() int64 {
	var last_id int64
	date := time.Now()
	dateConv := date.Format("1/2/2006")
	fmt.Println(date)
	result := db.QueryRow(`INSERT INTO pedido (data)
	 VALUES($1) RETURNING numero`, dateConv).Scan(&last_id)

	if result == sql.ErrNoRows {
		log.Panic("Error creating Pedido.")
		return 0
	}

	return last_id
}

// TODO - Debuggar o erro "violation of foreign key relashioship"
func CriaPedidoProduto(p PedidoProduto) int64 {
	var last_id int64
	result := db.QueryRow(`INSERT INTO pedidoproduto (codigo_produto, numero_pedido,quantidade,valorvenda)
	VALUES($1,$2,$3,$4) RETURNING codigo`, p.Codigo_produto, p.Numero_pedido, p.Quantidade, p.ValorVenda).Scan(&last_id)

	if result == sql.ErrNoRows {
		log.Panic("Error creating PedidoProduto.")
		return 0
	}

	return last_id
}
