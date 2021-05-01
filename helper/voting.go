package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// VoterData Struct
type VoterData struct {
	VoterId      int    `json:"voter_id"`
	PolicitianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}
type Voters []VoterData

// Connect to database
func connectVoting() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// DecodeVoter file json
func DecodeVoter() Voters {
	jsonFile, err := os.Open("json/voting.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Success open json file")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var voterData Voters
	err = json.Unmarshal(byteValue, &voterData)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(voterData)
	return voterData
}

//InsertVoter json to database
func InsertVoter(data Voters) {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success to database election")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	for _, value := range data {
		_, err = db.Exec("INSERT INTO votings VALUES (?, ?, ?, ?, ?, ?)", value.VoterId, value.PolicitianId, value.FirstName, value.LastName, value.Gender, value.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success insert voters to database")
	}
}

// Query show all voting
//func QueryShowAllVoting(){
//
//}

// Query data voting bergender male nama diawali huruf b
//func QueryShowMaleWithB(){
//
//}

func Voter() {
	// Decode Json
	tempVoter := DecodeVoter()
	fmt.Println(tempVoter)

	// insert json to database
	//InsertVoter(tempVoter)
}
