package models

import "api_postgres/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		//Trata no Handlers
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	//Quando a variável já foi inicializada, não é necessário o ":"
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}
