package helper

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost)/electionGo")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
