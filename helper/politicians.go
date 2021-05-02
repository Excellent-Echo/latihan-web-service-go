package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"sort"
)

// Politic data struct
type Politic struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}
type Politician []Politic

// Koneksi ke database MySQL
func connectPolitic() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/elections")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// DecodePolitic untuk decode file json politicians
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

//InsertPolitic untuk memasukkan semua data json politicians ke database MySQL
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

// QueryShowAllPolitic Query untuk menampilkan semua data politic
func QueryShowAllPolitic() {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData []Politic
	data, err := db.Query("SELECT * FROM politicians")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)

	for _, value := range politicData {
		fmt.Println(value)
	}
}

// QueryPopularVotingNY Query untuk menampilkan politicians yang populer di kawasasan NY
func QueryPopularVotingNY() {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData []Politic
	data, err := db.Query("SELECT * FROM politicians")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	max := politicData[0]
	for _, value := range politicData {
		if value.Location == "NY" {
			if value.GradeCurrent > max.GradeCurrent {
				max = value
			}
		}
	}
	fmt.Println(max)
}

// QueryPopularThreePopular Query untuk menampilkan 3 politicians terpopuler berdasarkan votingnya
func QueryPopularThreePopular() {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData []Politic
	data, err := db.Query("SELECT * FROM politicians")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)

	sort.Slice(politicData, func(i, j int) bool { return politicData[i].GradeCurrent < politicData[j].GradeCurrent })
	fmt.Println("first:", politicData[len(politicData)-1])
	fmt.Println("second:", politicData[len(politicData)-2])
	fmt.Println("three:", politicData[len(politicData)-3])
}

// QueryILLocation Query untuk menampilkan semua politician di location IL beserta votingnya
func QueryILLocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "IL"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// QueryNYLocation Query untuk menampilkan semua politician di location NY beserta votingnya
func QueryNYLocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "NY"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// QueryTXLocation Query untuk menampilkan semua politician di location TX beserta votingnya
func QueryTXLocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "TX"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// QueryLALocation Query untuk menampilkan semua politician di location LA beserta votingnya
func QueryLALocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "IL"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// QueryWALocation Query untuk menampilkan semua politician di location WA beserta votingnya
func QueryWALocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "WA"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// QueryFLLocation Query untuk menampilkan semua politician di location FL beserta votingnya
func QueryFLLocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "FL"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// QueryHILocation Query untuk menampilkan semua politician di location HI beserta votingnya
func QueryHILocation() Politician {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var politicData Politician

	var location = "HI"
	data, err := db.Query("SELECT * FROM politicians WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanData Politic
		if data.Scan(&satuanData.PoliticianId, &satuanData.Name, &satuanData.Party, &satuanData.Location, &satuanData.GradeCurrent); err != nil {
			fmt.Println(err.Error())
		}
		politicData = append(politicData, satuanData)
	}
	//fmt.Println(politicData)
	return politicData
}

// Function utama untuk eksekusi semua function bantuan
func Politicians() {
	// Function untuk decode file json politicians
	//tempPolitic := DecodePolitic()
	//fmt.Println(tempPolitic)

	// Function untuk memasukkan file json ke database MySQL
	//InsertPolitic(tempPolitic)

	// Function query untuk menampilkan semua data json politicians
	//QueryShowAllPolitic()

	// Function query untuk menampilkan politicians terpopuler di NY
	//QueryPopularVotingNY()

	// Function query untuk menampilkan 3 politicians terpopuler
	//QueryPopularThreePopular()
}
