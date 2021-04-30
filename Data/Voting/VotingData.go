package Voting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func VoterData() {
	var politiciansData Voters
	jsonFile, err := os.Open("Data/Voting/voting.json")
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file voting.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politiciansData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range politiciansData {
		fmt.Println(data.VoterId)
		fmt.Println(data.PoliticianId)
		fmt.Println(data.FirstName)
		fmt.Println(data.LastName)
		fmt.Println(data.Gender)
		fmt.Println(data.Age)
	}
}