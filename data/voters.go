package data

import (
	"electionGo/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Voter struct {
	ID           int    `json:"voter_id"`
	PoliticianID int    `json:"politician_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func DecodeVoters() {
	db, err := helper.Connect()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fileData, err := os.Open("data/voters.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, _ := ioutil.ReadAll(fileData)

	var voters []Voter

	json.Unmarshal(result, &voters)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS voters (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		politician_id INT NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		gender VARCHAR(10) NOT NULL,
		age INT NOT NULL,

		FOREIGN KEY (politician_id)
		REFERENCES politicians (id)
		ON UPDATE CASCADE
		)`)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, voter := range voters {

		_, err = db.Exec("INSERT INTO voters VALUES(?, ?, ?, ?, ?, ?)", voter.ID, voter.PoliticianID, voter.FirstName, voter.LastName, voter.Gender, voter.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success insert to database")
	}
}
