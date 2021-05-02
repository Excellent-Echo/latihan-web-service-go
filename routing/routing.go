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

// func votingsRoute() {
// 	// start query
// 	db, err := connect.Connect()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(
// 		`SELECT * FROM votings`,
// 	)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	var data []Voting

// 	for rows.Next() {
// 		var each = Voting{}
// 		var err = rows.Scan(&each.VoterID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		data = append(data, each)
// 	}

// 	if err = rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	// end query

// 	// web service API server
// 	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		if r.Method == "GET" {
// 			byteJson, err := json.Marshal(data)

// 			if err != nil {
// 				http.Error(w, "error internal server", http.StatusInternalServerError)
// 				return
// 			}

// 			w.Write(byteJson)
// 			return
// 		}

// 		http.Error(w, "", http.StatusBadRequest)
// 	})

// 	// port := "localhost:4444"

// 	// fmt.Println("server running on port", port)

// 	// http.ListenAndServe(port, nil)
// }

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

func Routing() {
	// votingsRoute()
	http.HandleFunc("/votings", votings)
	http.HandleFunc("/votings_male", votingsMale)

	port := "localhost:4444"

	fmt.Println("server running on port", port)

	http.ListenAndServe(port, nil)
}
