package routing

import (
	"encoding/json"
	"fmt"
	"latihan-web-service-go/connect"
	"net/http"
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

func votings(w http.ResponseWriter, r *http.Request) {
	// start query
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
		return
	}
	defer rows.Close()

	var data []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	// end query

	// web service API
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func votingsMale(w http.ResponseWriter, r *http.Request) {
	// start query
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
		return
	}
	defer rows.Close()

	var data []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	// end query

	// web service API
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func votingsFemale(w http.ResponseWriter, r *http.Request) {
	// start query
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
		return
	}
	defer rows.Close()

	var data []Voting

	for rows.Next() {
		var each = Voting{}
		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	// end query

	// web service API
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func politicians(w http.ResponseWriter, r *http.Request) {
	// start query
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM politicians`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var data []Politician

	for rows.Next() {
		var each = Politician{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	// end query

	// web service API
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func politiciansVoting(w http.ResponseWriter, r *http.Request) {
	// start query
	db, err := connect.Connect()
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

	var data []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	// end query

	// web service API
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func Routing() {
	// votingsRoute()
	http.HandleFunc("/votings", votings)
	http.HandleFunc("/votings_male", votingsMale)
	http.HandleFunc("/votings_female", votingsFemale)
	http.HandleFunc("/politicians", politicians)
	http.HandleFunc("/politicians_voting", politiciansVoting)

	port := "localhost:4444"
	fmt.Println("server running on port", port)
	http.ListenAndServe(port, nil)
}
