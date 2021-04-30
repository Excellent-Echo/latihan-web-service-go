package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Politic struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float64 `json:"grade_current"`
}

func connectPolitic() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func Politicians() {

	jsonFile, err := os.Open("json/politicians.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Success open json file")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var politicData []Politic
	err = json.Unmarshal(byteValue, &politicData)

	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(politicData)

	// Connect database
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success to database election")
	defer db.Close()

	//Insert All Data politicians to Database

	for _, value := range politicData {
		//fmt.Println("Id Politik :", value.PoliticianId)
		//fmt.Println("Name :", value.Name)
		//fmt.Println("Party :", value.Party)
		//fmt.Println("Location :", value.Location)
		//fmt.Println("Grade Current :", value.GradeCurrent)
		_, err = db.Exec("INSERT INTO politicians VALUES (?, ?, ?, ?, ?)", value.PoliticianId, value.Name, value.Party, value.Location, value.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("success insert to database")
	}

}
