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

func showVoting() (dataVote []Voting) {
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
	datas := showVoting()

	for _, data := range datas {
		fmt.Println(data.VotingID)
		fmt.Println(data.PoliticianID)
		fmt.Println(data.FirstName)
		fmt.Println(data.LastName)
		fmt.Println(data.Gender)
		fmt.Println(data.Age)
	}

}
