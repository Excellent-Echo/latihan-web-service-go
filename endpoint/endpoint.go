package endpoint

import (
	"encoding/json"
	"fmt"
	"latihan-web-service-go/helper"
	"net/http"
)

func Endpoint() {
	// Fungsi helper yang berisikan data json yang telah diquery
	dataVoters := helper.QueryShowAllVoting()
	maleVoters := helper.QueryShowMale()
	femaleVoters := helper.QueryShowFemale()
	dataPolitic := helper.QueryShowAllPolitic()
	dataPoliticAndVoting := helper.QueryShowAllPoliticVotingJoin()
	ilLocation := helper.QueryILLocation()
	nyLocation := helper.QueryNYLocation()
	txLocation := helper.QueryTXLocation()
	laLocation := helper.QueryLALocation()
	waLocation := helper.QueryWALocation()
	flLocation := helper.QueryFLLocation()
	hiLocation := helper.QueryHILocation()

	// Endpoint untuk menampilkan semua data voting
	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// Endpoint untuk menampilkan semua data voting gender male
	http.HandleFunc("/votings_male", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// Endpoint untuk menampilkan semua data voting gender female
	http.HandleFunc("/votings_female", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// Endpoint untuk menampilkan semua data politic
	http.HandleFunc("/politicians", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// Endpoint untuk menampilkan semua data politician beserta voting
	http.HandleFunc("/politicians_voting", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// Endpoint untuk menampilkan semua politician di location IL beserta votingnya
	http.HandleFunc("/politicians_il", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(ilLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// Endpoint untuk menampilkan semua politician di location NY beserta votingnya
	http.HandleFunc("/politicians_ny", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(nyLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// Endpoint untuk menampilkan semua politician di location TX beserta votingnya
	http.HandleFunc("/politicians_tx", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(txLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// Endpoint untuk menampilkan semua politician di location LA beserta votingnya
	http.HandleFunc("/politicians_la", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(laLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// Endpoint untuk menampilkan semua politician di location WA beserta votingnya
	http.HandleFunc("/politicians_wa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(waLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// Endpoint untuk menampilkan semua politician di location FL beserta votingnya
	http.HandleFunc("/politicians_fl", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(flLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// Endpoint untuk menampilkan semua politician di location HI beserta votingnya
	http.HandleFunc("/politicians_hi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(hiLocation)
			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "error not method GET", http.StatusBadRequest)
	})

	// untuk menyambungkan localhost dari local komputer
	port := "localhost:3000"
	fmt.Println("server running on port", port)
	http.ListenAndServe(port, nil)
}
