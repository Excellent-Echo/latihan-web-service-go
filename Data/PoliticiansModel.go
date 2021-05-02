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

type PoliticiansVoteDB struct {
	Politician_id int
	Name          string
	Party         string
	Location      string
	Grade_current float32
	Vote          int
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

//Get All data Politicians
func GetDataPoliticiansVoting() (dataPoliticiansDB []PoliticiansVoteDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voters AS v ON p.politician_id = v.politician_id GROUP BY p.politician_id;")
	if err != nil {
		fmt.Println(err.Error())
	}

	for data.Next() {

		var dat PoliticiansVoteDB

		if data.Scan(
			&dat.Politician_id,
			&dat.Name,
			&dat.Party,
			&dat.Location,
			&dat.Grade_current,
			&dat.Vote); err != nil {
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

//Get All Politician with highest SCORE
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

//Get All data Politicians From IL
func GetDataPoliticiansFromIL() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='IL'")
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

//Get All data Politicians From NY
func GetDataPoliticiansFromNY() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='NY'")
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

//Get All data Politicians From TX
func GetDataPoliticiansFromTX() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='TX'")
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

//Get All data Politicians From LA
func GetDataPoliticiansFromLA() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='LA'")
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

//Get All data Politicians From WA
func GetDataPoliticiansFromWA() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='WA'")
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

//Get All data Politicians From FL
func GetDataPoliticiansFromFL() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='FL'")
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

//Get All data Politicians From HI
func GetDataPoliticiansFromHI() (dataPoliticiansDB []PoliticiansDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM politicians WHERE location='HI'")
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
