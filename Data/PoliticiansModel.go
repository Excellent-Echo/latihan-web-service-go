package Data

import (
	"fmt"
)

type PoliticiansDB struct {
	Politician_id int
	Name          string
	Party         string
	Location      string
	Grade_current float32
}

//get data Politicians by ID
func GetDataPoliticianById(id int) (dataPoliticianDB PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE politician_id = ?", id)
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

		dataPoliticianDB = dat
	}
	return
}

//Get All data Politicians
func GetDataPoliticians() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
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

//Get All Politician with highest SCORE IN NY
func GetDataPoliticianWithHighestScoreOnNY() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='NY' ORDER BY grade_current DESC LIMIT 1")
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

//Get All Politician with highest SCORE IN NY
func GetDataPoliticiansHeadWithHighestScore() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians ORDER BY grade_current DESC LIMIT 3")
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
	db, err := Connect()
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
