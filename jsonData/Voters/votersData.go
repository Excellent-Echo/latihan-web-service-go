package Voters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Voters []struct {
	Voters_id      int    `json:"voters_id"`
	Politicians_id int    `json:"politician_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Gender         string `json:"gender"`
	Age            int    `json:"age"`
}

func VotersData() {
	var votersData Voters
	jsonFile, err := os.Open("jsonData/Voters/voting.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("sucsess get file Voters.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &votersData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range votersData {
		fmt.Println("voter_id", data.Voters_id)
		fmt.Println("politician_id", data.Politicians_id)
		fmt.Println("first_name", data.FirstName)
		fmt.Println("last_name", data.LastName)
		fmt.Println("gender", data.Gender)
		fmt.Println("age", data.Age)
	}

}
