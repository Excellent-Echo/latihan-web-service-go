package handler

import (
	"encoding/json"
	"latihan-web-service-go/helper"
	"net/http"
	"path"
	"text/template"
)

// RootHandler untuk menghandle jika ada url yang error
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

// VotingHandler Endpoint untuk menampilkan semua data voting
func VotingHandler(w http.ResponseWriter, r *http.Request) {
	dataVoters := helper.QueryShowAllVoting()

	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dataVoters)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(dataVoters)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// VotingMaleHandler Endpoint untuk menampilkan semua data voting gender male
func VotingMaleHandler(w http.ResponseWriter, r *http.Request) {
	maleVoters := helper.QueryShowMale()
	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, maleVoters)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(maleVoters)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// VotingFemaleHandler Endpoint untuk menampilkan semua data voting gender female
func VotingFemaleHandler(w http.ResponseWriter, r *http.Request) {
	femaleVoters := helper.QueryShowFemale()
	tmpl, err := template.ParseFiles(path.Join("views", "votings.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, femaleVoters)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(femaleVoters)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}
