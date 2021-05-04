package main

import (
	"fmt"
	"latihan-web-service-go/entity"
)

func main() {
	// insert politicians.json into DB
	politicians, err := entity.PoliticianData()
	if err != nil {
		fmt.Println("Error inserting into DB")
	}
	entity.SqlQueryPoliticians(politicians)

	// insert voting.json into DB
	votings, err := entity.VoterData()
	if err != nil {
		fmt.Println("Error inserting into DB")
	}
	entity.SqlQueryVoters(votings)
}
