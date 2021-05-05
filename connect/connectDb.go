package connect

import (
	"database/sql"
	"fmt"
)

func connet() (*sql.DB, error) {

	db, err := sql.Open("mysql", "Andryan:Celanadlm12!(localhost)/gopolitic")

	if err != nil {
		fmt.Println(err.Error())
	}

	return db, nil
}
