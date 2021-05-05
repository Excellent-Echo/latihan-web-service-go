package query

import (
	"electionGo/helper"
	"fmt"
)

type Politicians struct {
	vote_count    int
	name          string
	party         string
	location      string
	grade_current float32
}

type Voters struct {
	id            int
	politician_id int
	first_name    string
	last_name     string
	gender        string
	age           int
}

// displays politicians data and the number of votes
func Release2_1() {
	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT COUNT(voters.politician_id) AS vote_count, politicians.name, politicians.party, politicians.location, politicians.grade_current
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	GROUP BY politician_id`)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Politicians

	for rows.Next() {
		var each = Politicians{}
		var err = rows.Scan(&each.vote_count, &each.name, &each.party, &each.location, &each.grade_current)

		fmt.Println(each)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each)
	}

}

// Filter voter data based on male gender and their names begin with the letter B
func Release2_2() {
	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM voters WHERE gender='male' AND first_name LIKE 'B%'`)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Voters

	for rows.Next() {
		var each = Voters{}
		var err = rows.Scan(&each.id, &each.politician_id, &each.first_name, &each.last_name, &each.gender, &each.age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each)
	}

}

// Largest number of votes in NY
func Release2_3() {
	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT COUNT(voters.politician_id) AS vote_count, politicians.name, politicians.party, politicians.location, politicians.grade_current 
	FROM voters 
	INNER JOIN politicians ON voters.politician_id = politicians.id 
	WHERE location='NY'
	GROUP BY politician_id
	LIMIT 1`)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Politicians

	for rows.Next() {
		var each = Politicians{}
		var err = rows.Scan(&each.vote_count, &each.name, &each.party, &each.location, &each.grade_current)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each)
	}

}

//3 politicians with the most votes
func Release2_4() {
	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT COUNT(voters.politician_id) AS vote_count, politicians.name, politicians.party, politicians.location, politicians.grade_current 
	FROM voters 
	INNER JOIN politicians ON voters.politician_id = politicians.id 
	GROUP BY politician_id
	ORDER BY vote_count DESC
	LIMIT 3`)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Politicians

	for rows.Next() {
		var each = Politicians{}
		var err = rows.Scan(&each.vote_count, &each.name, &each.party, &each.location, &each.grade_current)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each)
	}

}
