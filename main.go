package main

import (
	"fmt"
	"latihan-web-service-go/entity"
)

func main() {


	//entity.VoterData()

	// insert politicians.json into DB
	politicians, err := entity.PoliticianData()
	if err != nil {
		fmt.Println("Error inserting into DB")
	}
	entity.SqlQuery(politicians)
}
