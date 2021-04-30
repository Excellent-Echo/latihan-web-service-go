package Data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type PoliticiansDB struct {
	Politician_id int
	Name          string
	Party         string
	Location      string
	Grade_current float32
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "Radika:!Satu234limA!@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}

func GetDataPoliticians() (dataPoliticiansDB []PoliticiansDB) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians")
	if err != nil {
		fmt.Println(err.Error())
	}

	for data.Next() {

		var dat PoliticiansDB

		if data.Scan(
			&dat.Politician_id,
			&dat.Name,
			&dat.Party,
			&dat.Location,
			&dat.Grade_current); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataPoliticiansDB = append(dataPoliticiansDB, dat)
	}
	return
}

//this function for initial post from json only do not use for others
func InitialPostDataPoliticians(id int, name string, party string, location string, grade float32) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//query
	_, err = db.Exec("INSERT INTO politicians VALUES (?,?,?,?,?)", id, name, party, location, grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Success initial insert data")
}
