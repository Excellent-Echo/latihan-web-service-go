package routing

import (
	"encoding/json"
	"fmt"
	"latihan-web-service-go/helper"
	"net/http"
)

func votings(w http.ResponseWriter, r *http.Request) {
	data := helper.AllVoters()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func votingsMale(w http.ResponseWriter, r *http.Request) {
	data := helper.MaleVoters()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func votingsFemale(w http.ResponseWriter, r *http.Request) {
	data := helper.FemaleVoters()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func politicians(w http.ResponseWriter, r *http.Request) {
	data := helper.AllPoliticianDataQuery()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func politiciansVoting(w http.ResponseWriter, r *http.Request) {
	data := helper.AllPoliticianWithVoteQuery()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func ILpoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.ILPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func NYpoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.NYPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func TXpoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.TXPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func LApoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.LAPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func WApoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.WAPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func FLpoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.FLPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func HIpoliticians(w http.ResponseWriter, r *http.Request) {
	data := helper.HIPoliticians()

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func Routing() {
	http.HandleFunc("/votings", votings)
	http.HandleFunc("/votings_male", votingsMale)
	http.HandleFunc("/votings_female", votingsFemale)
	http.HandleFunc("/politicians", politicians)
	http.HandleFunc("/politicians_voting", politiciansVoting)
	http.HandleFunc("/politicians_il", ILpoliticians)
	http.HandleFunc("/politicians_ny", NYpoliticians)
	http.HandleFunc("/politicians_tx", TXpoliticians)
	http.HandleFunc("/politicians_la", LApoliticians)
	http.HandleFunc("/politicians_wa", WApoliticians)
	http.HandleFunc("/politicians_fl", FLpoliticians)
	http.HandleFunc("/politicians_hi", HIpoliticians)

	port := "localhost:4444"
	fmt.Println("server running on port", port)
	http.ListenAndServe(port, nil)
}
