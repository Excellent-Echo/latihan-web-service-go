package entity

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-service-go/config"
	"os"
)

type Voters []struct {
	VoterId 		int 	`json:"voter_id"`
	PoliticianId 	int 	`json:"policitian_id"`
	FirstName 		string 	`json:"first_name"`
	LastName 		string 	`json:"last_name"`
	Gender 			string 	`json:"gender"`
	Age				int 	`json:"age"`
}

func VoterData() (Voters, error) {
	var votersData Voters

	jsonFile, err := os.Open("data/voting.json")
	if err != nil{
		return Voters{}, err
	}
	fmt.Println("success get file voting.json")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &votersData)
	if err != nil{
		return Voters{}, err
	}
	return votersData, nil
}

func SqlQueryVoters(votersData Voters)  {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	for _, data := range votersData {
		_, err = db.Exec("INSERT INTO votings " +
			"VALUES (?, ?, ?, ?, ?, ?)",
			data.VoterId,
			data.PoliticianId,
			data.FirstName,
			data.LastName,
			data.Gender,
			data.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("insert data success")
	}
}