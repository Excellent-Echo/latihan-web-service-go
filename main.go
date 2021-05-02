package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type PoliticianVoting struct {
	Voter_id         int
	Politician_id_fk int
	First_name       string
	Last_name        string
	Gender           string
	Age              int
	Politician_id    int
	Name             string
	Party            string
	Location         string
	Grade_current    float32
}
type Politician struct {
	Politician_id int     `json:"politician_id"`
	Name          string  `json:"name"`
	Party         string  `json:"party"`
	Location      string  `json:"location"`
	Grade_current float32 `json:"grade_current"`
}

type Voting struct {
	Voter_id      int    `json:"voter_id"`
	Politician_id int    `json:"policitian_id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
}

type PoliticianWithSumVoting struct {
	Politician_id int
	Name          string
	Party         string
	Location      string
	Grade_current float32
	Voting        int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/voting_app")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func importJson(file string) []byte {
	fileJson := file + ".json"
	res, err := os.Open(fileJson)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("success baca file ", fileJson)

	defer res.Close()

	byteValue, _ := ioutil.ReadAll(res)
	return byteValue
}

func getDataPolitician(query string) (rowPoliticians []Politician) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.Politician_id, &politician.Name, &politician.Party, &politician.Location, &politician.Grade_current)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rowPoliticians = append(rowPoliticians, politician)
	}
	return
}

func getDataVoting(query string) (rowVotings []Voting) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var voting Voting
		err := rows.Scan(&voting.Voter_id, &voting.Politician_id, &voting.First_name, &voting.Last_name, &voting.Gender, &voting.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rowVotings = append(rowVotings, voting)
	}
	return
}

func getPoliticianVoting(query string) (rowsPoliticianVoting []PoliticianVoting) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var politicianVoting PoliticianVoting
		err := rows.Scan(&politicianVoting.Voter_id, &politicianVoting.Politician_id_fk, &politicianVoting.First_name, &politicianVoting.Last_name, &politicianVoting.Gender, &politicianVoting.Age, &politicianVoting.Politician_id, &politicianVoting.Name, &politicianVoting.Party, &politicianVoting.Location, &politicianVoting.Grade_current)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rowsPoliticianVoting = append(rowsPoliticianVoting, politicianVoting)
	}
	return
}

func getPoliticianWithSumVoting(query string) (rowsPoliticianWithSumVoting []PoliticianWithSumVoting) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var politicianWithSumVoting PoliticianWithSumVoting
		err := rows.Scan(&politicianWithSumVoting.Politician_id, &politicianWithSumVoting.Name, &politicianWithSumVoting.Party, &politicianWithSumVoting.Location, &politicianWithSumVoting.Grade_current, &politicianWithSumVoting.Voting)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rowsPoliticianWithSumVoting = append(rowsPoliticianWithSumVoting, politicianWithSumVoting)
	}
	return
}

func main() {
	//REALESE 1 : IMPORT JSON & INSERT TO DATABASE

	// import data politicians
	byteValuePolitician := importJson("politicians")
	var politicianData []Politician
	json.Unmarshal(byteValuePolitician, &politicianData)

	//import data voting
	byteValueVoting := importJson("voting")
	var votingData []Voting
	json.Unmarshal(byteValueVoting, &votingData)

	//connect database
	// db, err := connect()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	//insert data politicians
	// for _, data := range politicianData {
	// 	_, err := db.Exec("INSERT INTO politicians (politician_id,name,party,location,grade_current) VALUES (?,?,?,?,?)", data.Politician_id, data.Name, data.Party, data.Location, data.Grade_current)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// }
	// fmt.Println("success insert data politician")

	//insert data voting
	// for _, data := range votingData {
	// 	_, err := db.Exec("INSERT INTO votings (voter_id, politician_id, first_name, last_name, gender, age) VALUES (?,?,?,?,?,?)", data.Voter_id, data.Politician_id, data.First_name, data.Last_name, data.Gender, data.Age)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// }
	// fmt.Println("success insert data voting")

	//REALESE 2

	//show politician and voting
	politicianVoting := getPoliticianVoting("select * from votings join politicians on votings.politician_id = politicians.politician_id")
	fmt.Println(politicianVoting)
	fmt.Println("=====================================")

	//show voting where gender male and first letter name B
	votingsMaleFirstLetterB := getDataVoting("select * from votings where gender='male' AND first_name LIKE 'B%'")
	fmt.Println(votingsMaleFirstLetterB)
	fmt.Println("=====================================")

	//show top 1 politician voting with location NY
	top1PoliticianInNY := getPoliticianWithSumVoting("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN votings AS v ON p.politician_id = v.politician_id WHERE location = 'NY' GROUP BY p.politician_id ORDER BY voting DESC LIMIT 1;")
	fmt.Println(top1PoliticianInNY)
	fmt.Println("=====================================")

	//show top 3 politician voting
	top3Politician := getPoliticianWithSumVoting("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN votings AS v ON p.politician_id = v.politician_id GROUP BY p.politician_id ORDER BY voting DESC LIMIT 3;")
	fmt.Println(top3Politician)
	fmt.Println("=====================================")

	//RELEASE 3
	http.HandleFunc("/politician", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("politician.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := getDataPolitician("select * from politicians")
		t.Execute(w, data)
	})

	port := "localhost:8000"

	fmt.Println("service running on port", port)

	http.ListenAndServe(port, nil)
}
