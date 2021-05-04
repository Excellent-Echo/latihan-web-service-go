package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/latihan_web_service_go")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success connect to database")

	return db, nil
}