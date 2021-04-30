package Data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "Radika:paswword@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
