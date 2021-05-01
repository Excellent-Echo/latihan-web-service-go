package main

import (
	"fmt"
	"latihan-web-service-go/Data"
)

func main() {
	//initial get data from json to database, just activate the function in beginning
	//Data.PoliticianData()
	//Data.VoterData()

	// Get Data ALL politician from database
	// datas := Data.GetDataPoliticians()

	//Get Data politician by id
	//datas := Data.GetDataPoliticianById(1)

	//Get Data Male Voter with Bname
	// datas := Data.GetDataMaleVotersWithBname()
	// for _, data := range datas {
	// 	fmt.Println(data.Voter_id)
	// 	fmt.Println(data.Politician_id)
	// 	fmt.Println(data.First_name)
	// 	fmt.Println(data.Last_name)
	// 	fmt.Println(data.Gender)
	// 	fmt.Println(data.Age)
	// }

	// Get Politican with highest score on NY
	//datas := Data.GetDataPoliticianWithHighestScoreOnNY()

	datas := Data.GetDataPoliticiansHeadWithHighestScore()
	for _, data := range datas {
		fmt.Println(data.Politician_id)
		fmt.Println(data.Name)
		fmt.Println(data.Party)
		fmt.Println(data.Location)
		fmt.Println(data.Grade_current)
	}
}
