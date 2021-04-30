package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type VoterData struct {
	VoterId      int    `json:"voter_id"`
	PolicitianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func connectVoting() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func Voter() {

	jsonFile, err := os.Open("json/voting.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Success open json file")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var voterData []VoterData
	err = json.Unmarshal(byteValue, &voterData)

	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(voterData)

	// Connect database
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success to database election")
	defer db.Close()

	for _, value := range voterData {
		fmt.Println("Id Voter :", value.VoterId)
		fmt.Println("Id Politic :", value.PolicitianId)
		fmt.Println("Firstname :", value.FirstName)
		fmt.Println("Lastname :", value.LastName)
		fmt.Println("Gender :", value.Gender)
		fmt.Println("Age :", value.Age)

		_, err = db.Exec("INSERT INTO votings VALUES (?, ?, ?, ?, ?, ?)", value.VoterId, value.PolicitianId, value.FirstName, value.LastName, value.Gender, value.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("success insert to database")
	}

}
