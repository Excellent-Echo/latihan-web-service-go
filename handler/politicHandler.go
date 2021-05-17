package handler

import (
	"encoding/json"
	"latihan-web-service-go/helper"
	"net/http"
	"path"
	"text/template"
)

// PoliticHandler Endpoint untuk menampilkan semua data politic
func PoliticHandler(w http.ResponseWriter, r *http.Request) {
	dataPolitic := helper.QueryShowAllPolitic()
	tmpl, err := template.ParseFiles(path.Join("views", "politicians.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dataPolitic)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(dataPolitic)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticAndVoting Endpoint untuk menampilkan semua data politician beserta voting
func PoliticAndVoting(w http.ResponseWriter, r *http.Request) {
	dataPoliticAndVoting := helper.QueryShowAllPoliticVotingJoin()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dataPoliticAndVoting)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(dataPoliticAndVoting)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticILLocation Endpoint untuk menampilkan semua politician di location IL beserta votingnya
func PoliticILLocation(w http.ResponseWriter, r *http.Request) {
	ilLocation := helper.QueryILLocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, ilLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(ilLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticNYLocation Endpoint untuk menampilkan semua politician di location NY beserta votingnya
func PoliticNYLocation(w http.ResponseWriter, r *http.Request) {
	nyLocation := helper.QueryNYLocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nyLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(nyLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticTXLocation Endpoint untuk menampilkan semua politician di location TX beserta votingnya
func PoliticTXLocation(w http.ResponseWriter, r *http.Request) {
	txLocation := helper.QueryTXLocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, txLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(txLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticLALocation Endpoint untuk menampilkan semua politician di location LA beserta votingnya
func PoliticLALocation(w http.ResponseWriter, r *http.Request) {
	laLocation := helper.QueryLALocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, laLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(laLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticWALocation Endpoint untuk menampilkan semua politician di location WA beserta votingnya
func PoliticWALocation(w http.ResponseWriter, r *http.Request) {
	waLocation := helper.QueryWALocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, waLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(waLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticFLLocation Endpoint untuk menampilkan semua politician di location FL beserta votingnya
func PoliticFLLocation(w http.ResponseWriter, r *http.Request) {
	flLocation := helper.QueryFLLocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, flLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(flLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}

// PoliticHILocation Endpoint untuk menampilkan semua politician di location HI beserta votingnya
func PoliticHILocation(w http.ResponseWriter, r *http.Request) {
	hiLocation := helper.QueryHILocation()
	tmpl, err := template.ParseFiles(path.Join("views", "politicAndVoters.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, hiLocation)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//result, err := json.Marshal(hiLocation)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}
		//w.Write(result)
		return
	}
	http.Error(w, "error not method GET", http.StatusBadRequest)
}
