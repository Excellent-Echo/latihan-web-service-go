function main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type DataBuild struct {
	Data []Politic `Politic`
}

type Politic struct {
	IdPolitician 	int		`json : "politician_id"`
	Name 			string	`json : "name"`
	Party 			string	`json : "party"`
	Location 		string	`json : "location"`
	Grade 			float32 `json : "grade_current"`
}

func main() {
	politician, err := os.Open("politicians.json")

	if err != nil {
		fmt.Println(err.Error)
	}

	fmt.Println("success open JSON")

	defer politician.Close()

	byteValue, _ := ioutil.ReadAll(politician)

	var dataResult DataBuild

	json.Unmarshal(byteValue, &dataResult)
}