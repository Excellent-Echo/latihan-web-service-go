package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Politician struct {
	PoliticianID   int     `json:"politician_id"`
	NamePolitician string  `json:"name"`
	Party          string  `json:"party"`
	Location       string  `json:"location"`
	GradeCurrent   float32 `json:"grade_current"`
}

type Politicians []Politician

type Voting struct {
	VotingID     int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type VotingData []Voting

func showVotingbyGender() (dataVote []Voting) {
	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()

	data, err := db.Query("select * from voting where gender='male' and first_name like 'B%'")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for data.Next() {

		var dat Voting

		if data.Scan(
			&dat.VotingID,
			&dat.PoliticianID,
			&dat.FirstName,
			&dat.LastName,
			&dat.Gender,
			&dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataVote = append(dataVote, dat)
	}
	return

}

func showPoliticianTop() (dataPolitician []Politician) {
	db, err := connect()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success get db")

	defer db.Close()

	data, err := db.Query("select * from politicians where location='NY' ORDER BY grade_current LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for data.Next() {

		var dat Politician

		if data.Scan(
			&dat.PoliticianID,
			&dat.NamePolitician,
			&dat.Party,
			&dat.Location,
			&dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataPolitician = append(dataPolitician, dat)
	}
	return

}

func connect() (*sql.DB, error) {

	// "username-mysql:password-mysql@tcp()/nama-database"
	db, err := sql.Open("mysql", "root:@tcp(localhost)/electionus")

	if err != nil {
		panic(err)
	}
	return db, nil
}

func main() {
	// getPoliticians()
	datas := showVotingbyGender()

	for _, data := range datas {
		fmt.Println(data.VotingID)
		fmt.Println(data.PoliticianID)
		fmt.Println(data.FirstName)
		fmt.Println(data.LastName)
		fmt.Println(data.Gender)
		fmt.Println(data.Age)
	}

	datasPol := showPoliticianTop()

	for _, data := range datasPol {
		fmt.Println(data.PoliticianID)
		fmt.Println(data.NamePolitician)
		fmt.Println(data.Party)
		fmt.Println(data.Location)
		fmt.Println(data.GradeCurrent)
	}

}
