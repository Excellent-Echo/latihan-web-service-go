package main

import (
	"database/sql"
	"fmt"

	// "encoding/json"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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

func main() {
	//REALESE 1 : IMPORT JSON & INSERT TO DATABASE

	//import data politicians
	// byteValuePolitician := importJson("politicians")
	// var politicianData []Politician
	// json.Unmarshal(byteValuePolitician, &politicianData)

	// //import data voting
	// byteValueVoting := importJson("voting")
	// var votingData []Voting
	// json.Unmarshal(byteValueVoting, &votingData)

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
	allPoliticians := getDataPolitician("select * from politicians")
	allVotings := getDataVoting("select * from votings")
	fmt.Println(allPoliticians)
	fmt.Println(allVotings)

	votingsWithFilter := getDataVoting("select * from votings where gender='male' AND first_name LIKE 'B%'")
	fmt.Println(votingsWithFilter)

	politiciansWithFilter2 := getDataPolitician("select * from politicians where location='NY' order by grade_current DESC limit 1")
	fmt.Println(politiciansWithFilter2)

	politiciansWithFilter3 := getDataPolitician("select * from politicians order by grade_current DESC limit 3")
	fmt.Println(politiciansWithFilter3)
}
