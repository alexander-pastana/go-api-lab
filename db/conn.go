package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//cria a conexão com o bd

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = 1234
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	//Variável que guarda a string com os dados de acesso ao bd
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", host, port, user, password, dbname)

	//Função que valida os argumentos e inicializa a estrutura de conexão
	db, err := sql.Open("postgres", psqInfo)

	if err != nil {
		panic(err)
	}

	//Para testar se o banco está realmente acessível. Ping precisa do sql.Open
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to " + dbname)

	return db, nil

}