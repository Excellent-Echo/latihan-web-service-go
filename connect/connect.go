package connect

import "database/sql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:danangaji@tcp(localhost)/db_voting")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}