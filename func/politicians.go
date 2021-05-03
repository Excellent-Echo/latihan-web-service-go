package func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Politicians []struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

func decodePoliticians() {
	var politiciansData Politicians

	jsonFile, err := os.Open("politicians.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file politicians.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politiciansData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range politiciansData {
		_, err = db.Exec("INSERT INTO politicians VALUES(?, ?, ?, ?, ?)", val.PoliticianId, val.Name, val.Party, val.Location, val.GradeCurrent)
		fmt.Println("successfully inserted data")
	}
}