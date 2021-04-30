package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Politician struct {
	Politician_id int     `json:"politician_id"`
	Name          string  `json:"name"`
	Party         string  `json:"party"`
	Location      string  `json:"location"`
	Grade_current float32 `json:"grade_current"`
}

type Voting struct {
	Voter_id      int    `json:"voter_id"`
	Politician_id int    `json:"politician_id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Gender        string `json:"gender"`
	Age           string `json:"age"`
}

func connect() {
	// db, err :=sql.Open("mysql","root")
}

func main() {
	politicians, err := os.Open("politicians.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success baca file json")

	defer politicians.Close()

	byteValue, _ := ioutil.ReadAll(politicians)
	var result []Politician
	json.Unmarshal(byteValue, &result)
	fmt.Println(result)
}
