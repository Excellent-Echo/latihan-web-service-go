package entity

import (
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
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politiciansData)
	if err != nil {
		return Politicians{}, err
	}
	return politiciansData, nil

	//for _, data := range politiciansData {
	//	fmt.Println(data.PoliticianId)
	//	fmt.Println(data.Name)
	//	fmt.Println(data.Party)
	//	fmt.Println(data.Location)
	//	fmt.Println(data.GradeCurrent)
	//}
}

func SqlQuery(politiciansData Politicians)  {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	for _, data := range politiciansData {
		_, err = db.Exec("INSERT INTO politicians VALUES (?, ?, ?, ?, ?)",
			data.PoliticianId, data.Name, data.Party, data.Location, data.GradeCurrent)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("insert data success")
	}
}