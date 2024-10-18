package model

import (
	"database/sql"
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

func CreateTables(db *sql.DB) *sql.Row {

	query := `
	CREATE TABLE IF NOT EXISTS pedido(
		numero INT NOT NULL PRIMARY KEY,
		data DATE
	);
	
	CREATE TABLE IF NOT EXISTS produto(
		codigo INT NOT NULL PRIMARY KEY,
		descricao VARCHAR (255),
		preco DOUBLE (2)
	);

	CREATE TABLE IF NOT EXISTS  departamento(
		codigo INT NOT NULL	PRIMARY KEY,
		codigo_produto INT NOT NULL,
		descricao VARCHAR (255),
		FOREIGN KEY (codigo_produto) REFERENCES Produto(codigo)
	);
	CREATE IF NOT EXISTS TABLE pedidoProduto(
		codigo INT NOT NULL PRIMARY KEY,
		codigo_produto INT NOT NULL,
		numero_pedido INT NOT NULL,
		quantidade INT,
		valorVenda DOUBLE(2),
		FOREIGN KEY (codigo_produto) REFERENCES Produto(codigo)
		FOREIGN KEY (numero_pedido) REFERENCES Pedido(numero)
	);
	`
	row := db.QueryRow(query)

	if row.Err() != nil {
		log.Fatal("Error creating the tables", row)
	}

	return row
}
