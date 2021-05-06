package query

import (
	"electionGo/helper"
	"encoding/json"
	"fmt"
	"net/http"
)

type Politician struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
	VoteCount    int     `json:"vote_count"`
}

type Voter struct {
	ID           int    `json:"id"`
	PoliticianID int    `json:"politician_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

// displays politicians data and the number of votes
func Release2_1() {
	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	GROUP BY politician_id`)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var each Politician
		var err = rows.Scan(&each.ID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.VoteCount)

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

	var result []Voter

	for rows.Next() {
		var each Voter
		var err = rows.Scan(&each.ID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

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

	rows, err := db.Query(`SELECT COUNT(voters.politician_id) AS vote_count, politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current
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

	var result []Politician

	for rows.Next() {
		var each Politician
		var err = rows.Scan(&each.VoteCount, &each.ID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent)

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

	rows, err := db.Query(`SELECT COUNT(voters.politician_id) AS vote_count, politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current
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

	var result []Politician

	for rows.Next() {
		var each Politician
		var err = rows.Scan(&each.VoteCount, &each.ID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent)

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

func GetAllVoters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM voters`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Voter

	for rows.Next() {
		var voter Voter
		err := rows.Scan(&voter.ID, &voter.PoliticianID, &voter.FirstName, &voter.LastName, &voter.Gender, &voter.Age)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, voter)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllMaleVoters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM voters WHERE gender='male'`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Voter

	for rows.Next() {
		var voter Voter
		err := rows.Scan(&voter.ID, &voter.PoliticianID, &voter.FirstName, &voter.LastName, &voter.Gender, &voter.Age)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, voter)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllFemaleVoters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM voters WHERE gender='female'`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Voter

	for rows.Next() {
		var voter Voter
		err := rows.Scan(&voter.ID, &voter.PoliticianID, &voter.FirstName, &voter.LastName, &voter.Gender, &voter.Age)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, voter)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticians(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM politicians`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansVoting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansIL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='IL'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansNY(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='NY'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansTX(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='TX'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansLA(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='LA'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansWA(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='WA'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansFL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='FL'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllPoliticiansHI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := helper.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT politicians.id, politicians.name, politicians.party, politicians.location, politicians.grade_current, COUNT(voters.politician_id) AS vote_count
	FROM voters
	INNER JOIN politicians ON voters.politician_id = politicians.id
	WHERE location='HI'
	GROUP BY politician_id`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []Politician

	for rows.Next() {
		var politician Politician
		err := rows.Scan(&politician.ID, &politician.Name, &politician.Party, &politician.Location, &politician.GradeCurrent, &politician.VoteCount)

		if err != nil {
			panic(err.Error())
		}

		result = append(result, politician)
	}

	json.NewEncoder(w).Encode(result)
}
