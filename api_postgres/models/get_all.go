package models

import "api_postgres/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	print("a")
	if err != nil {
		return
	}
	defer conn.Close()
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}
	//Percorre a lista e joga cada valor na lista
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}
