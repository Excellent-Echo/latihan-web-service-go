package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Politician struct {
	PoliticianID int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

type Politicians []Politician

type PoliticianWithTotalVotes struct {
	Politician
	TotalVotes int
}

type Voting struct {
	VoterID      int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type Votes []Voting

func decodePoliticianData(file string) Politicians {
	var politicians Politicians

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &politicians)
	if err != nil {
		fmt.Println(err.Error())
	}

	return politicians
}

func insertPoliticianDataToDb(p Politicians) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range p {
		_, err = db.Exec("INSERT INTO politicians (name, party, location, grade_current) VALUES (?, ?, ?, ?)", val.Name, val.Party, val.Location, val.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("insert data succeed")
}

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
	db, err := connect()
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

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:marwanajunolii@tcp(localhost)/election")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}

func allPoliticianQuery() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		GROUP BY p.politician_id`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var p []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

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
		fmt.Println("Politician ID:", each.PoliticianID)
		fmt.Println("Name:", each.Name)
		fmt.Println("Party:", each.Party)
		fmt.Println("Location:", each.Location)
		fmt.Println("Grade Current:", each.GradeCurrent)
		fmt.Println("Total Votes:", each.TotalVotes)
	}
}

func voterMaleQuery() {
	db, err := connect()
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

func popularPoliticianNYQuery() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT politician_id, name, party, location, grade_current, MAX(total) total
		FROM (SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
					 COUNT(v.politician_id) as total
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = 'NY'
		GROUP BY p.politician_id) cp`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var p []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

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
		fmt.Println("Politician ID:", each.PoliticianID)
		fmt.Println("Name:", each.Name)
		fmt.Println("Party:", each.Party)
		fmt.Println("Location:", each.Location)
		fmt.Println("Grade Current:", each.GradeCurrent)
		fmt.Println("Total Votes:", each.TotalVotes)
	}
}

func threePopularPoliticianQuery() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		GROUP BY p.politician_id
		ORDER BY total DESC
		LIMIT 3`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var p []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

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
		fmt.Println("Politician ID:", each.PoliticianID)
		fmt.Println("Name:", each.Name)
		fmt.Println("Party:", each.Party)
		fmt.Println("Location:", each.Location)
		fmt.Println("Grade Current:", each.GradeCurrent)
		fmt.Println("Total Votes:", each.TotalVotes)
	}
}

func votingsRoute(v Votes) {
	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			byteJson, err := json.Marshal(v)

			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}

			w.Write(byteJson)
			return

		}

		http.Error(w, "error not method GET", http.StatusInternalServerError)
	})

	port := "localhost:4444"

	fmt.Println("server running on port", port)

	http.ListenAndServe(port, nil)
}

func main() {
	// p := decodePoliticianData("politicians.json")
	v := decodeVotingData("voting.json")

	// insertPoliticianDataToDb(p)
	// insertVotingDataToDb(v)

	// allPoliticianQuery()
	// voterMaleQuery()
	// popularPoliticianNYQuery()
	// threePopularPoliticianQuery()
	votingsRoute(v)
}
