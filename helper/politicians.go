package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Politic struct
type Politic struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float64 `json:"grade_current"`
}
type Politician []Politic

// Connect to database
func connectPolitic() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// DecodePolitic file json
func DecodePolitic() Politician {
	jsonFile, err := os.Open("json/politicians.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Success open json file")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var politicData Politician
	err = json.Unmarshal(byteValue, &politicData)

	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(politicData)

	return politicData
}

//InsertPolitic json to database
func InsertPolitic(data Politician) {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	for _, value := range data {
		_, err = db.Exec("INSERT INTO politicians VALUES (?, ?, ?, ?, ?)", value.PoliticianId, value.Name, value.Party, value.Location, value.GradeCurrent)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success insert politic to database")
	}
}

//func QueryShowAllPolitic(){
//	db, err := connectPolitic()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//
//		}
//	}(db)
//
//	var name = "Aaron Schock"
//	data, err := db.Query("SELECT * FROM politicians WHERE name = ?",name)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	for data.Next() {
//		var satuanData Politic
//		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent);
//			err != nil {
//			fmt.Println(err.Error())
//			return
//		}
//		politicData = append(politicData, satuanData)
//	}
//	fmt.Println("test", politicData)
//}

func Politicians() {
	// Decode Json
	tempPolitic := DecodePolitic()
	fmt.Println(tempPolitic)

	// insert json to database
	//InsertPolitic(tempPolitic)

	// Query all politicians data
	//QueryShowAllPolitic()
}
