package WebService

import (
	"fmt"
	"latihan-web-service-go/function"
	"net/http"
	"path"
	"text/template"
)

func Service() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "index.html"), path.Join("webService/pages/", "layout.html"))

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, nil)
	})

	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "votings.html"), path.Join("webService/pages/", "layout.html"))

		// data := map[string]string{}
		var data = function.GetVotings()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// fmt.Println(data)
		// for i := 0; i < len(datas); i++ {
		// 	data["voter_id"] = "<h2>" + strconv.Itoa(datas[i].VoterId) + "</h2>"
		// 	// data["voter_id"] = ""
		// 	data["politician_id"] = strconv.Itoa(datas[i].PoliticianId)
		// 	// first_name := datas[i].FirstName
		// 	// last_name

		// }

		t.Execute(w, data)
	})

	http.HandleFunc("/votings_male", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "votingMale.html"), path.Join("webService/pages/", "layout.html"))

		data := function.QueryVotingMale()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)

	})

	http.HandleFunc("/votings_female", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "votingFemale.html"), path.Join("webService/pages/", "layout.html"))

		data := function.QueryVotingFemale()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)

	})

	http.HandleFunc("/politicians", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politicians.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := function.GetPoliticians()

		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_voting", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansVotings.html"), path.Join("webService/pages/", "layout.html"))

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := function.QueryPoliticiansWithVoting()

		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_il", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocIL.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := function.QueryPoliticians1stVotingLocIL()

		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_ny", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocNY.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := function.QueryPoliticians1stVotingLocNY()

		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_tx", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocTX.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := function.QueryPoliticians1stVotingLocTX()
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_la", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocLA.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := function.QueryPoliticians1stVotingLocLA()
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_wa", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocWA.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := function.QueryPoliticians1stVotingLocWA()
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_fl", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocFL.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := function.QueryPoliticians1stVotingLocFL()
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_hi", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join("webService/pages/", "politiciansLocFL.html"), path.Join("webService/pages/", "layout.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := function.QueryPoliticians1stVotingLocHI()
		t.Execute(w, data)
	})

	port := "localhost:8080"
	fmt.Println("Starting Web Service at", port)
	http.ListenAndServe(port, nil)
}
