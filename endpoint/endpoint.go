package endpoint

import (
	"encoding/json"
	"fmt"
	"latihan-web-service-go/helper"
	"net/http"
)

func Endpoint() {
	data := helper.DecodeVoter()
	male := helper.EndpointQueryShowMale()
	female := helper.EndpointQueryShowFemale()
	dataPolitic := helper.DecodePolitic()

	// Endpoint untuk menampilkan semua data voting
	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(data)
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
			result, err := json.Marshal(male)
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
			result, err := json.Marshal(female)
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

	port := "localhost:3000"
	fmt.Println("server running on port", port)
	http.ListenAndServe(port, nil)
}
