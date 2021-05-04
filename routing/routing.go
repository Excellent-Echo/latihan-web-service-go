package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webservice/config"
	"webservice/entity"
)

func ShowVoting(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	dataSQL, err := db.Query("select * from  voting")

	if err != nil {
		fmt.Println("error query SQL")
		return
	}

	var datas []entity.Voting

	for dataSQL.Next() {
		var each = entity.Voting{}

		if err := dataSQL.Scan(&each.VotingID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age); err != nil {
			log.Fatal(err.Error())
			return
		}

		datas = append(datas, each)
	}
	defer db.Close()

	r.Header.Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// excecute yg bener
		byteDatas, err := json.Marshal(datas)

		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		w.Write(byteDatas)
		return
	}

	// kalau bukan method get tampilkan error
	http.Error(w, "error method not GET", http.StatusBadRequest)
}

func ShowVotingbyMale(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	dataSQL, err := db.Query("select * from  voting where gender='male'")

	if err != nil {
		fmt.Println("error query SQL")
		return
	}

	var datas []entity.Voting

	for dataSQL.Next() {
		var each = entity.Voting{}

		if err := dataSQL.Scan(&each.VotingID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age); err != nil {
			log.Fatal(err.Error())
			return
		}

		datas = append(datas, each)
	}
	defer db.Close()

	r.Header.Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// excecute yg bener
		byteDatas, err := json.Marshal(datas)

		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		w.Write(byteDatas)
		return
	}

	// kalau bukan method get tampilkan error
	http.Error(w, "error method not GET", http.StatusBadRequest)
}

func ShowVotingbyFemale(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	dataSQL, err := db.Query("select * from  voting where gender='female'")

	if err != nil {
		fmt.Println("error query SQL")
		return
	}

	var datas []entity.Voting

	for dataSQL.Next() {
		var each = entity.Voting{}

		if err := dataSQL.Scan(&each.VotingID, &each.PoliticianID, &each.FirstName, &each.LastName, &each.Gender, &each.Age); err != nil {
			log.Fatal(err.Error())
			return
		}

		datas = append(datas, each)
	}
	defer db.Close()

	r.Header.Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// excecute yg bener
		byteDatas, err := json.Marshal(datas)

		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		w.Write(byteDatas)
		return
	}

	// kalau bukan method get tampilkan error
	http.Error(w, "error method not GET", http.StatusBadRequest)
}

func ShowPolitician(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	dataSQL, err := db.Query("select * from politicians")

	if err != nil {
		fmt.Println("error query SQL")
		return
	}

	var datas []entity.Politician

	for dataSQL.Next() {
		var each = entity.Politician{}

		if err := dataSQL.Scan(&each.PoliticianID, &each.NamePolitician, &each.Party, &each.Location, &each.GradeCurrent); err != nil {
			log.Fatal(err.Error())
			return
		}

		datas = append(datas, each)
	}
	defer db.Close()

	r.Header.Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// excecute yg bener
		byteDatas, err := json.Marshal(datas)

		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		w.Write(byteDatas)
		return
	}

	// kalau bukan method get tampilkan error
	http.Error(w, "error method not GET", http.StatusBadRequest)
}
