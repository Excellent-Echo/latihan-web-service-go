package webserver

import (
	"fmt"
	"html/template"
	"latihan-web-service-go/helper"
	"net/http"
)

type Voting struct {
	VoterID      int    `json:"voter_id"`
	PoliticianID int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type Votes []Voting

func votingsView() {
	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
		data := helper.AllVoters()

		for _, val := range data {
			fmt.Println(val)
		}

		var t, err = template.ParseFiles("webserver/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})
	port := "localhost:8080"

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(port, nil)
}

func WebServer() {
	votingsView()

}
