package entity

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-service-go/config"
	"os"
)

type Politicians []struct {
	PoliticianId 	int 	`json:"politician_id"`
	Name 			string 	`json:"name"`
	Party 			string 	`json:"party"`
	Location 		string 	`json:"location"`
	GradeCurrent 	float32	`json:"grade_current"`
}

func PoliticianData() (Politicians, error) {
	var politiciansData Politicians

	jsonFile, err := os.Open("data/politicians.json")
	if err != nil {
		return Politicians{}, err
	}
	fmt.Println("success get file politicians.json")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politiciansData)
	if err != nil {
		return Politicians{}, err
	}
	return politiciansData, nil
}

func SqlQueryPoliticians(politiciansData Politicians)  {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	for _, data := range politiciansData {
		_, err = db.Exec("INSERT INTO politicians " +
			"VALUES (?, ?, ?, ?, ?)",
			data.PoliticianId,
			data.Name,
			data.Party,
			data.Location,
			data.GradeCurrent)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("insert data success")
	}
}