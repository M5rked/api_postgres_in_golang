package db

import (
	"api_postgres/configs"
	"database/sql"
	"fmt"

	//O underline força o GO a não excluir a dependência
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	conn, err := sql.Open("postgres", sc)
	//Aqui tá escrito "se a conexão não funcionar"
	if err != nil {
		panic(err)
	}
	//Pra conferir se está estabelecida a conexão
	err = conn.Ping()
	return conn, err
}
