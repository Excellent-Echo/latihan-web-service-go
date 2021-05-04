package main

import (
	"fmt"
	"html/template"
	"latihan-web-service-go/Data"
	"net/http"
	"path"
)

func main() {
	//initial get data from json to database, just activate the function in beginning
	//Data.PoliticianData()
	//Data.VoterData()

	//Route Votings
	http.HandleFunc("/votings", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("View", "ViewVoting.html"), path.Join("View", "Layout.html"))

		datas := Data.GetDataVoters()

		if err != nil {
			http.Error(rw, "error rendering html", http.StatusInternalServerError)
			return
		}

		t.Execute(rw, datas)
	})

	//Route Voting Male
	http.HandleFunc("/votings_male", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		datas := Data.GetDataMaleVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	//Route Voting Female
	http.HandleFunc("/votings_female", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		datas := Data.GetDataFemaleVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	//Route Politicians
	http.HandleFunc("/politicians", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		datas := Data.GetDataPoliticians()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, datas)
	})

	//Route Politicians Voting
	http.HandleFunc("/politicians_voting", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansVoting()
		//votingData := Data.GetDataVoters()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, data := range politiciansData {
			fmt.Println(data.Vote)
		}
		t.Execute(rw, politiciansData)
	})

	//Route Politicians IL
	http.HandleFunc("/politicians_il", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromIL()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	//Route Politicians NY
	http.HandleFunc("/politicians_ny", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromNY()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	//Route Politicians TX
	http.HandleFunc("/politicians_tx", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromTX()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	//Route Politicians LA
	http.HandleFunc("/politicians_la", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromLA()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	//Route Politicians WA
	http.HandleFunc("/politicians_wa", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromWA()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	//Route Politicians FL
	http.HandleFunc("/politicians_fl", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromFL()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	//Route Politicians HI
	http.HandleFunc("/politicians_hi", func(rw http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("View/ViewVoting.html")

		politiciansData := Data.GetDataPoliticiansFromHI()
		//votingData := Data.GetDataVoters()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(rw, politiciansData)
	})

	port := "localhost:8080"
	fmt.Println("Server Running on port :", port)
	http.ListenAndServe(port, nil)
}
