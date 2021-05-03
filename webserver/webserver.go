package webserver

import (
	"html/template"
	"latihan-web-service-go/helper"
	"net/http"
	"path"
)

type Voting struct {
	VoterID      int
	PoliticianID int
	FirstName    string
	LastName     string
	Gender       string
	Age          int
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// kembalikan error
		http.Error(w, "404 error page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "root.html"), path.Join("views", "layout.html"))
	// w.Write([]byte("ini adalah root route"))
	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}

func AllVotingHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.AllVoters()

	// var list []Voting

	// for _, val := range data {
	// 	list = []Voting{
	// 		{val.VoterID, val.PoliticianID, val.FirstName, val.LastName, val.Gender, val.Age},
	// 	}
	// }

	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
