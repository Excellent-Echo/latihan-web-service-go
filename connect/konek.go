package connect

import "database/sql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Masukinaja123@tcp(localhost)/pemilu")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
