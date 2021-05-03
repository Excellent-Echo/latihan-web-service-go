package func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Voters []struct {
	VoterID      int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func decodeVoting() {
	var votersData Voters

	jsonFile, err := os.Open("voting.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file voting.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &votersData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range votersData {
		_, err = db.Exec("insert into voting values(?,?,?,?,?,?)", val.VoterId, val.PoliticianId, val.FirstName, val.LastName, val.Gender, val.Age)
		fmt.Println("successfully inserted data")
	}
}	