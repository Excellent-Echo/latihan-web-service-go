package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	// "username-mysql:password-mysql@tcp()/nama-database"
	db, err := sql.Open("mysql", "root:@tcp(localhost)/electionus")

	if err != nil {
		panic(err)
	}
	return db
}
