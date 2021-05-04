package query

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Politician struct {
	PoliticianID   int     `json:"politician_id"`
	NamePolitician string  `json:"name"`
	Party          string  `json:"party"`
	Location       string  `json:"location"`
	GradeCurrent   float32 `json:"grade_current"`
}

type Politicians []Politician

type Voting struct {
	VotingID     int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type VotingData []Voting

func insertPoliticians(file string) {
	var politicians Politicians
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("success get file politicians.json")
	// defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politicians)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// for _, data := range politicianData {
	// 	fmt.Println("politician_id : ", data.PoliticianID)
	// 	fmt.Println("name : ", data.NamePolitician)
	// 	fmt.Println("party : ", data.Party)
	// 	fmt.Println("location : ", data.Location)
	// 	fmt.Println("grade_current : ", data.GradeCurrent)
	// }

	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()
	for _, data := range politicians {
		_, err = db.Exec("INSERT INTO politicians (name, party, location, grade_current) VALUES (?, ?, ?, ?)", data.NamePolitician, data.Party, data.Location, data.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func insertVoting(file string) {
	var dataVot VotingData
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("success get file voting.json")
	// defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &dataVot)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// for _, data := range dataVot {
	// 	fmt.Println("voter_id : ", data.VotingID)
	// 	fmt.Println("politician_id : ", data.PoliticianID)
	// 	fmt.Println("first_name : ", data.FirstName)
	// 	fmt.Println("last_name : ", data.LastName)
	// 	fmt.Println("gender : ", data.Gender)
	// 	fmt.Println("age : ", data.Age)
	// }

	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()
	for _, data := range dataVot {
		_, err = db.Exec("INSERT INTO voting (politician_id, first_name, last_name, gender, age) VALUES (?, ?, ?, ?, ?)", data.PoliticianID, data.FirstName, data.LastName, data.Gender, data.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}

func showVotingbyGender() (dataVote []Voting) {
	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()

	data, err := db.Query("select * from voting where gender='male' and first_name like 'B%'")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for data.Next() {

		var dat Voting

		if data.Scan(
			&dat.VotingID,
			&dat.PoliticianID,
			&dat.FirstName,
			&dat.LastName,
			&dat.Gender,
			&dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataVote = append(dataVote, dat)
	}
	return

}

func showPoliticianTopbyLocation() (dataPolitician []Politician) {
	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()

	data, err := db.Query("select * from politicians where location='NY' ORDER BY grade_current LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for data.Next() {

		var dat Politician

		if data.Scan(
			&dat.PoliticianID,
			&dat.NamePolitician,
			&dat.Party,
			&dat.Location,
			&dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataPolitician = append(dataPolitician, dat)
	}
	return

}

func showPoliticianTop() (dataPolitician []Politician) {
	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()

	data, err := db.Query("select * from politicians ORDER BY grade_current LIMIT 3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for data.Next() {

		var dat Politician

		if data.Scan(
			&dat.PoliticianID,
			&dat.NamePolitician,
			&dat.Party,
			&dat.Location,
			&dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataPolitician = append(dataPolitician, dat)
	}
	return

}

func connect() (*sql.DB, error) {

	// "username-mysql:password-mysql@tcp()/nama-database"
	db, err := sql.Open("mysql", "root:@tcp(localhost)/electionus")

	if err != nil {
		panic(err)
	}
	return db, nil
}
