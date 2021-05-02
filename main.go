package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Data struct {
	Politic []Politician
}

type Politician []struct {
	Politician int    `json:"politician_id"`
	Name       string `json:"name"`
	Party      string `json:"party"`
	Location   string `json:"location"`
	Grade      int    `json:"grade_current"`
}

func main() {
	jsonFilePolitician, err := os.Open("politicians.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	byteValue, err := ioutil.ReadAll(jsonFilePolitician)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var dataPolitic Data

	json.Unmarshal([]byte(byteValue), &dataPolitic)

	for i := 0; i < len(dataPolitic.Politic); i++ {
		fmt.Println("politian_id: " + strconv.Itoa(dataPolitic.Politic[i].Politician))
		fmt.Println("name: " + dataPolitic.Politic[i].Name)
		fmt.Println("party: " + dataPolitic.Politic[i].Party)
		fmt.Println("location: " + dataPolitic.Politic[i].Location)
		fmt.Println("grade_current: " + strconv.Itoa(dataPolitic.Politic[i].Grade))
	}
}
