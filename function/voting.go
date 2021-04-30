package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Voting []struct {
	VoterId      int    `json:"voter_id"`
	PoliticianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func VotingFunc() {
	jsonVoting, err := os.Open("voting.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer jsonVoting.Close()

	byteVoting, _ := ioutil.ReadAll(jsonVoting)

	var voting Voting
	json.Unmarshal([]byte(byteVoting), &voting)
	// fmt.Println("Data Voting JSON")
	// for _, val := range voting {
	// 	fmt.Println(val)
}
