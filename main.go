package main

import (
	"fmt"
	"latihan-web-service-go/entity"
)

func main() {


	//entity.VoterData()
	politicians, err := entity.PoliticianData()
	if err != nil {
		fmt.Println("eror")
	}
	//fmt.Println(politicians)
	entity.SqlQuery(politicians)
}
