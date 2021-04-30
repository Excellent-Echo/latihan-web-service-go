package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Politician struct {
	PoliticianID int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

type Politicians []Politician

type Voting struct {
	VoterID      int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type Votes []Voting

func decodePoliticianData(file string) Politicians {
	var politicians Politicians

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &politicians)
	if err != nil {
		fmt.Println(err.Error())
	}

	return politicians
}

func insertPoliticianDataToDb(p Politicians) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range p {
		_, err = db.Exec("INSERT INTO politicians (name, party, location, grade_current) VALUES (?, ?, ?, ?)", val.Name, val.Party, val.Location, val.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("insert data succeed")
}

func decodeVotingData(file string) Votes {
	var votes Votes

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &votes)
	if err != nil {
		fmt.Println(err.Error())
	}

	return votes
}

func insertVotingDataToDb(v Votes) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range v {
		_, err = db.Exec("INSERT INTO votings (politician_id, first_name, last_name, gender, age) VALUES (?, ?, ?, ?, ?)", val.PoliticianID, val.FirstName, val.LastName, val.Gender, val.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	fmt.Println("insert data succeed")
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:marwanajunolii@tcp(localhost)/election")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}

func main() {
	// p := decodePoliticianData("politicians.json")
	v := decodeVotingData("voting.json")

	// insertPoliticianDataToDb(p)
	insertVotingDataToDb(v)
}
