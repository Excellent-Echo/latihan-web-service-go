package helper

import (
	"encoding/json"
	"fmt"
	"latihan-web-service-go/connect"
	"latihan-web-service-go/entity"
	"log"
	"net/http"
)

func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dataSQL, err := db.Query("SELECT * FROM politicians")
	if err != nil {
		fmt.Println("error query sql")
		return
	}

	var datas []entity.Politician

	for dataSQL.Next() {
		var each = entity.Politician{}

		if err = dataSQL.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent); err != nil {
			log.Fatal(err.Error())
			return
		}

		datas = append(datas, each)
	}

	defer db.Close()

	r.Header.Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// execute benar
		byteDatas, err := json.Marshal(datas)
		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		w.Write(byteDatas)
		return
	}

	//kalau bukan method get
	http.Error(w, "error method request", http.StatusBadRequest)

}
