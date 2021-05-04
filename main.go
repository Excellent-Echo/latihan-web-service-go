package main

import (
	"fmt"
	"net/http"
	"webservice/routing"
)

// type Politician struct {
// 	PoliticianID   int     `json:"politician_id"`
// 	NamePolitician string  `json:"name"`
// 	Party          string  `json:"party"`
// 	Location       string  `json:"location"`
// 	GradeCurrent   float32 `json:"grade_current"`
// }

// type Politicians []Politician

// type Voting struct {
// 	VotingID     int    `json:"voter_id"`
// 	PoliticianID int    `json:"policitian_id"`
// 	FirstName    string `json:"first_name"`
// 	LastName     string `json:"last_name"`
// 	Gender       string `json:"gender"`
// 	Age          int    `json:"age"`
// }

// type VotingData []Voting

func main() {
	http.HandleFunc("/voting", routing.ShowVoting)

	port := "localhost:3000"

	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
