package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-service-go/connect"
	"os"
)

type Voting struct {
	VoterID      int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type Votes []Voting

func decodeVotingData(file string) Votes {
	var votes Votes

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &votes)
	if err != nil {
		fmt.Println(err.Error())
	}

	return votes
}

func insertVotingDataToDb(v Votes) {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range v {
		_, err = db.Exec("INSERT INTO votings (politician_id, first_name, last_name, gender, age) VALUES (?, ?, ?, ?, ?)", val.PoliticianID, val.FirstName, val.LastName, val.Gender, val.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	fmt.Println("insert data succeed")
}

func voterMaleQuery() {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM votings
		WHERE gender = 'male' AND first_name LIKE 'B%'`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var p []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		p = append(p, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range p {
		fmt.Println("Voter ID:", each.VoterID)
		fmt.Println("First Name:", each.FirstName)
		fmt.Println("Last Name:", each.LastName)
		fmt.Println("Gender:", each.Gender)
		fmt.Println("Age:", each.Age)
		fmt.Println("Vote For Politician ID:", each.PoliticianID)
	}
}

// query for end point /votings
func AllVoters() Votes {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM votings`,
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var data []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	return data
}

// query for end point /votings_male
func MaleVoters() Votes {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM votings
		WHERE gender = 'male'`,
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var data []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	return data
}

// query for end point /votings_female
func FemaleVoters() Votes {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM votings
		WHERE gender = 'female'`,
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var data []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	return data
}

func Voter() {
	// v := decodeVotingData("json/voting.json")

	// insertVotingDataToDb(v)
	voterMaleQuery()
}
