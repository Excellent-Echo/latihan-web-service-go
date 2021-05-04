package Data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// POLITICIANS DATA

type Voters []struct {
	Voter_id      int    `json:"voter_id"`
	Politician_id int    `json:"policitian_id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
}

func VoterData() {

	var votersData Voters

	//get data json
	jsonFile, err := os.Open("Data/voting.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file voting.json")
	defer jsonFile.Close()

	//read data json
	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &votersData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//store to database
	for _, data := range votersData {
		InitialPostDataVoters(
			data.Voter_id,
			data.Politician_id,
			data.First_name,
			data.Last_name,
			data.Gender,
			data.Age,
		)
	}

	// for _, data := range votersData {
	// 	fmt.Println(data.Voter_id)
	// 	fmt.Println(data.Politician_id)
	// 	fmt.Println(data.First_name)
	// 	fmt.Println(data.Last_name)
	// 	fmt.Println(data.Gender)
	// 	fmt.Println(data.Age)
	// }
}
