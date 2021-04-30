package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	PoliticianID int    `json:"politician_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type Votes []Voting

func decodeJsonData(file string) {
	var politicians Politicians
	var votes Votes

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if file == "politicians.json" {
		err = json.Unmarshal(byteValue, &politicians)
	} else if file == "voting.json" {
		err = json.Unmarshal(byteValue, &votes)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, val := range politicians {
		fmt.Println("Politician ID: ", val.PoliticianID)
		fmt.Println("Name: ", val.Name)
		fmt.Println("Party: ", val.Party)
		fmt.Println("Location: ", val.Location)
		fmt.Println("Grade Current: ", val.GradeCurrent)
	}

	for _, val := range votes {
		fmt.Println("Voter ID: ", val.VoterID)
		fmt.Println("Politician ID: ", val.PoliticianID)
		fmt.Println("First Name: ", val.FirstName)
		fmt.Println("Last Name: ", val.LastName)
		fmt.Println("Gender: ", val.Gender)
		fmt.Println("Age: ", val.Age)
	}
}

func main() {
	decodeJsonData("politicians.json")
	decodeJsonData("voting.json")
}
