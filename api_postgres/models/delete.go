package models

import "api_postgres/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos where id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
