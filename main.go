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

	//Route Votings
	http.HandleFunc("/votings", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("viewVoting.html")

		datas := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	//Route Voting Male
	http.HandleFunc("/votings_male", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("viewVoting.html")

		datas := Data.GetDataMaleVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	//Route Voting Female
	http.HandleFunc("/votings_female", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("viewVoting.html")

		datas := Data.GetDataFemaleVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	http.HandleFunc("/politicians", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("viewVoting.html")

		datas := Data.GetDataPoliticians()

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
