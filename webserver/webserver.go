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

type Politician struct {
	PoliticianID int
	Name         string
	Party        string
	Location     string
	GradeCurrent float32
}

type PoliticianWithTotalVotes struct {
	Politician
	TotalVotes int
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

	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func MaleVotingHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.MaleVoters()

	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func FemaleVotingHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.FemaleVoters()

	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func AllPoliticianHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.AllPoliticianDataQuery()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func AllPoliticianWithVoteHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.AllPoliticianWithVoteQuery()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func ILPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.ILPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func NYPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.NYPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func TXPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.TXPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func LAPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.LAPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func WAPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.WAPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func FLPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.FLPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func HIPoliticiansHandler(w http.ResponseWriter, r *http.Request) {
	data := helper.HIPoliticians()

	tmpl, err := template.ParseFiles(path.Join("views", "politicians_vote.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
