package decodeJson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Politician struct {
	PoliticianID   int     `json:"politician_id"`
	NamePolitician string  `json:"name"`
	Party          string  `json:"party"`
	Location       string  `json:"location"`
	GradeCurrent   float32 `json:"grade_current"`
}

type Politicians []Politician

type Voting struct {
	VotingID     int    `json:"voter_id"`
	PoliticianID int    `json:"politician_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type VotingData []Voting

func getPoliticians(file string) {
	var politicianData Politicians
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file politicians.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politicianData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range politicianData {
		fmt.Println("politician_id : ", data.PoliticianID)
		fmt.Println("name : ", data.NamePolitician)
		fmt.Println("party : ", data.Party)
		fmt.Println("location : ", data.Location)
		fmt.Println("grade_current : ", data.GradeCurrent)
	}
}

func getVoting(file string) {
	var dataVot VotingData
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file voting.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &dataVot)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range dataVot {
		fmt.Println("voter_id : ", data.VotingID)
		fmt.Println("politician_id : ", data.PoliticianID)
		fmt.Println("first_name : ", data.FirstName)
		fmt.Println("last_name : ", data.LastName)
		fmt.Println("gender : ", data.Gender)
		fmt.Println("age : ", data.Age)
	}
}
