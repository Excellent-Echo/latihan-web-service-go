package data

import (
	"electionGo/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Politician struct {
	ID           int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

// type Politician struct {
// 	ID           int
// 	Name         string
// 	Party        string
// 	Location     string
// 	GradeCurrent float32
// }

func DecodePoliticians() {
	db, err := helper.Connect()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fileData, err := os.Open("data/politicians.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, _ := ioutil.ReadAll(fileData)

	var politicians []Politician

	json.Unmarshal(result, &politicians)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS politicians (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		party VARCHAR(100) NOT NULL,
		location VARCHAR(2) NOT NULL,
		grade_current DOUBLE NOT NULL)`)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, politician := range politicians {
		// var data = fmt.Sprintf("(%d, '%s', '%s', '%s', %g)", politician.ID, politician.Name, politician.Party, politician.Location, politician.GradeCurrent)

		// insert += data + ", "

		_, err = db.Exec("INSERT INTO politicians VALUES(?, ?, ?, ?, ?)", politician.ID, politician.Name, politician.Party, politician.Location, politician.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success insert to database")
	}
	// var execute = fmt.Sprintf("INSERT INTO politicians VALUES %s", insert)
}
