package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Politician []struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

type Voting []struct {
	VoterId      int    `json:"voter_id"`
	PoliticianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func main() {
	//Politicians JSON
	jsonPoliticians, err := os.Open("politicians.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success open jsonPoliticians")

	defer jsonPoliticians.Close()

	bytePocisians, _ := ioutil.ReadAll(jsonPoliticians)

	var politician Politician

	json.Unmarshal([]byte(bytePocisians), &politician)

	fmt.Println("Data Politician JSON")
	for _, val := range politician {
		fmt.Println(val)
	}
	fmt.Println("\n")

	// Voting JSON
	jsonVoting, err := os.Open("voting.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer jsonVoting.Close()

	byteVoting, _ := ioutil.ReadAll(jsonVoting)

	var voting Voting
	json.Unmarshal([]byte(byteVoting), &voting)
	fmt.Println("Data Voting JSON")
	for _, val := range voting {
		fmt.Println(val)
	}
	// fmt.Println("\n")
}
