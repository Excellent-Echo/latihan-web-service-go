package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-service-go/connect"
	"latihan-web-service-go/entity"
	"os"
)

func decodeVotingData(file string) entity.Votes {
	var votes entity.Votes

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

func insertVotingDataToDb(v entity.Votes) {
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

	var p []entity.Voting

	for rows.Next() {
		var each = entity.Voting{}
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
func AllVoters() entity.Votes {
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

	var data []entity.Voting

	for rows.Next() {
		var each = entity.Voting{}
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
func MaleVoters() entity.Votes {
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

	var data []entity.Voting

	for rows.Next() {
		var each = entity.Voting{}
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
func FemaleVoters() entity.Votes {
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

	var data []entity.Voting

	for rows.Next() {
		var each = entity.Voting{}
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
