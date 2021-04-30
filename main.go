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

func main() {
	var politicians Politicians

	jsonFile, err := os.Open("politicians.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Success open jsonFile")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &politicians)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(politicians)
}
