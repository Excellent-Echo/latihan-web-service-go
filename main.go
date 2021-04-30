package main

import (
	"fmt"
	"latihan-web-service-go/Data"
)

func main() {
	//initial get datafrom json to database, just activate the function in beginning
	//Data.PoliticianData()

	datas := Data.GetDataPoliticians()
	for _, data := range datas {
		fmt.Println(data.Politician_id)
		fmt.Println(data.Name)
		fmt.Println(data.Party)
		fmt.Println(data.Location)
		fmt.Println(data.Grade_current)
	}
}
