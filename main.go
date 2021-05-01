package main

import (
	"fmt"
	"html/template"
	"latihan-web-service-go/Data"
	"net/http"
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
	// datas := Data.GetDataMaleVotersWithBname()git

	// Get Politican with highest score on NY
	//datas := Data.GetDataPoliticianWithHighestScoreOnNY()

	// Get 3 Politicians highest Score
	//datas := Data.GetDataPoliticiansHeadWithHighestScore()

	http.HandleFunc("/votings", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("viewVoting.html")

		datas := Data.GetDataVoters()
		// for _, data := range datas {
		// 	fmt.Println(data.Voter_id)
		// 	fmt.Println(data.Politician_id)
		// 	fmt.Println(data.First_name)
		// 	fmt.Println(data.Last_name)
		// 	fmt.Println(data.Gender)
		// 	fmt.Println(data.Age)
		// }
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	http.HandleFunc("/politicians", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("viewVoting.html")

		datas := Data.GetDataPoliticians()
		// for _, data := range datas {
		// 	fmt.Println(data.Voter_id)
		// 	fmt.Println(data.Politician_id)
		// 	fmt.Println(data.First_name)
		// 	fmt.Println(data.Last_name)
		// 	fmt.Println(data.Gender)
		// 	fmt.Println(data.Age)
		// }
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	port := "localhost:8080"
	fmt.Println("Server Running on port :", port)
	http.ListenAndServe(port, nil)
}
