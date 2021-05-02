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

// PoliticVoterDataJoin data struct
type PoliticVoterDataJoin struct {
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
	VoterId      int     `json:"voter_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Gender       string  `json:"gender"`
	Age          int     `json:"age"`
}
type PoliticVotersJoin []PoliticVoterDataJoin

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

// QueryShowAllPoliticVotingJoin Query untuk menampilkan semua data Polticians dan Voting yang sudah di join table
func QueryShowAllPoliticVotingJoin() PoliticVotersJoin {
	db, err := connectVoting()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin
	data, err := db.Query("SELECT * FROM joindata")
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryShowAllPolitic Query untuk menampilkan semua data politic
func QueryShowAllPolitic() Politician {
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
	data, err := db.Query("SELECT * FROM politicians")
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
	fmt.Println(politicData)
	return politicData
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
func QueryILLocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "IL"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryNYLocation Query untuk menampilkan semua politician di location NY beserta votingnya
func QueryNYLocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "NY"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryTXLocation Query untuk menampilkan semua politician di location TX beserta votingnya
func QueryTXLocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "TX"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryLALocation Query untuk menampilkan semua politician di location LA beserta votingnya
func QueryLALocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "LA"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryWALocation Query untuk menampilkan semua politician di location WA beserta votingnya
func QueryWALocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "WA"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryFLLocation Query untuk menampilkan semua politician di location FL beserta votingnya
func QueryFLLocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "FL"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// QueryHILocation Query untuk menampilkan semua politician di location HI beserta votingnya
func QueryHILocation() PoliticVotersJoin {
	db, err := connectPolitic()
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var PoliticVotingDataJoin PoliticVotersJoin

	var location = "HI"
	data, err := db.Query("SELECT * FROM joindata WHERE location = ?", location)
	if err != nil {
		fmt.Println(err.Error())
	}
	for data.Next() {
		var satuanDataJoin PoliticVoterDataJoin
		if data.Scan(&satuanDataJoin.VoterId, &satuanDataJoin.Name, &satuanDataJoin.Party, &satuanDataJoin.Location, &satuanDataJoin.GradeCurrent, &satuanDataJoin.FirstName, &satuanDataJoin.LastName, &satuanDataJoin.Gender, &satuanDataJoin.Age); err != nil {
			fmt.Println(err.Error())
		}
		PoliticVotingDataJoin = append(PoliticVotingDataJoin, satuanDataJoin)
	}
	//fmt.Println(PoliticVotingDataJoin)
	return PoliticVotingDataJoin
}

// Politicians Function utama untuk eksekusi semua function bantuan
func Politicians() {
	// Function untuk decode file json politicians
	//tempPolitic := DecodePolitic()
	//fmt.Println(tempPolitic)

	// Function untuk memasukkan file json ke database MySQL
	//InsertPolitic(tempPolitic)

	// Function query untuk menampilkan semua data json politicians dan voting join table
	QueryShowAllPoliticVotingJoin()

	// Function query untuk menampilkan semua data json politicians
	//QueryShowAllPolitic()

	// Function query untuk menampilkan politicians terpopuler di NY
	//QueryPopularVotingNY()

	// Function query untuk menampilkan 3 politicians terpopuler
	//QueryPopularThreePopular()
}
