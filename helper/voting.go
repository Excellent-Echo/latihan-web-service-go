package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// VoterData data struct
type VoterData struct {
	VoterId      int    `json:"voter_id"`
	PolicitianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}
type Voters []VoterData

// Koneksi ke database MySQL
func connectVoting() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// DecodeVoter untuk decode file json voting
func DecodeVoter() Voters {
	jsonFile, err := os.Open("json/voting.json")
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
	var voterData Voters
	err = json.Unmarshal(byteValue, &voterData)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(voterData)
	return voterData
}

//InsertVoter untuk memasukkan semua data json voter ke database MySQL
func InsertVoter(data Voters) {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success to database election")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	for _, value := range data {
		_, err = db.Exec("INSERT INTO votings VALUES (?, ?, ?, ?, ?, ?)", value.VoterId, value.PolicitianId, value.FirstName, value.LastName, value.Gender, value.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success insert voters to database")
	}
}

// QueryShowAllVoting Query untuk menampilkan semua data voting
func QueryShowAllVoting() {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var votingData []VoterData
	data, err := db.Query("SELECT * FROM votings")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for data.Next() {
		var satuanData VoterData
		if data.Scan(&satuanData.VoterId, &satuanData.PolicitianId, &satuanData.FirstName, &satuanData.LastName, &satuanData.Gender, &satuanData.Age); err != nil {
			fmt.Println(err.Error())
			return
		}
		votingData = append(votingData, satuanData)
	}
	//fmt.Println(votingData)

	for _, value := range votingData {
		fmt.Println(value)
	}
}

// QueryShowMaleWithB Query untuk menampilkan data voters laki-laki yang namanya diawali huruf B
func QueryShowMaleWithB() {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var votingData []VoterData

	var gender = "male"
	data, err := db.Query("SELECT * FROM votings WHERE gender = ?", gender)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for data.Next() {
		var satuanData VoterData
		if data.Scan(&satuanData.VoterId, &satuanData.PolicitianId, &satuanData.FirstName, &satuanData.LastName, &satuanData.Gender, &satuanData.Age); err != nil {
			fmt.Println(err.Error())
			return
		}
		votingData = append(votingData, satuanData)
	}
	fmt.Println(votingData)

	for _, value := range votingData {
		//fmt.Println(value)
		if string(value.FirstName[0]) == "B" {
			fmt.Println(value)
		}
	}
}

// QueryShowMale Query untuk menampilkan semua data voting yang berjenis kelamin laki-laki
func QueryShowMale() Voters {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var votingData Voters
	var gender = "male"
	data, err := db.Query("SELECT * FROM votings WHERE gender = ?", gender)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData VoterData
		if data.Scan(&satuanData.VoterId, &satuanData.PolicitianId, &satuanData.FirstName, &satuanData.LastName, &satuanData.Gender, &satuanData.Age); err != nil {
			fmt.Println(err.Error())
		}
		votingData = append(votingData, satuanData)
	}
	//fmt.Println(votingData)
	return votingData
}

// QueryShowFemale Query untuk menampilkan semua data voting berjenis kelamin perempuan
func QueryShowFemale() Voters {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var votingData Voters
	var gender = "female"
	data, err := db.Query("SELECT * FROM votings WHERE gender = ?", gender)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData VoterData
		if data.Scan(&satuanData.VoterId, &satuanData.PolicitianId, &satuanData.FirstName, &satuanData.LastName, &satuanData.Gender, &satuanData.Age); err != nil {
			fmt.Println(err.Error())
		}
		votingData = append(votingData, satuanData)
	}
	//fmt.Println(votingData)
	return votingData
}

// Function utama untuk eksekusi semua function bantuan
func Voter() {
	// Function untuk decode file json voting
	//tempVoter := DecodeVoter()
	//fmt.Println(tempVoter)

	// Function untuk memasukkan file json ke database MySQL
	//InsertVoter(tempVoter)

	// Function query untuk menampilkan semua data json voting
	//QueryShowAllVoting()

	// Function query untuk menampilkan semua data voter laki-laki yang namanya diawali huruf B
	//QueryShowMaleWithB()
}
