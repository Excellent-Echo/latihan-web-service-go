package main

import "latihan-web-service-go/Data"

func main() {
	//initial get data from json to database, just activate the function in beginning
	//Data.PoliticianData()
	Data.VoterData()

	// Get Data ALL politician from database
	// datas := Data.GetDataPoliticians()
	// for _, data := range datas {
	// 	fmt.Println(data.Politician_id)
	// 	fmt.Println(data.Name)
	// 	fmt.Println(data.Party)
	// 	fmt.Println(data.Location)
	// 	fmt.Println(data.Grade_current)
	// }

	//Get Data politician by id
	//datas := Data.GetDataPoliticianById(1)

}
