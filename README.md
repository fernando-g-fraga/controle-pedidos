# Controle Pedidos

This project is a Go-based application designed to manage product orders. It uses the Echo framework for the web server and PostgreSQL for the database.
Prerequisites

    Go 1.23.2 or later
    PostgreSQL 17
    Git

Dependencies

The project uses the following Go modules:

    Echo v4.12.0: A high-performance, extensible, minimalist web framework for Go.
    lib/pq v1.10.9: PostgreSQL driver for Go.
    Additional indirect dependencies are handled automatically by Go's dependency management.

All dependencies are managed in the go.mod file.

### Todo List
Table Pedidos:
- [x]  Create
- [x]  Delete
- [ ]  Read ID
- [ ]  Update

Table Departamento:
- [ ]  Read

Advanced Stuff: 
- [ ]  JointTable pedidoProduto
- [ ]  Ordering and Filters

Installation

1. Clone the Repository

git clone https://github.com/fernandogfaga/controle-pedidos.git
cd controle-pedidos

2. Install Dependencies

Run the following command to install the required Go modules:
go mod tidy

3. Set Up PostgreSQL Database

Ensure that you have PostgreSQL 17 installed and running.
Create a new database:
CREATE DATABASE controle_pedidos;
Update the database connection settings in your Go project as needed.

4. Run the Application

You can start the application with:

go run main.go
