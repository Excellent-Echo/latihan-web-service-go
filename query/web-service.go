package query

// import (
// 	"electionGo/helper"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type Voter struct {
// 	ID           int    `json:"id"`
// 	PoliticianID int    `json:"politician_id"`
// 	FirstName    string `json:"first_name"`
// 	LastName     string `json:"last_name"`
// 	Gender       string `json:"gender"`
// 	Age          int    `json:"age"`
// }

// func GetAllVoters(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	db, err := helper.Connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(`SELECT * FROM voters`)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer rows.Close()

// 	var result []Voter

// 	for rows.Next() {
// 		var voter Voter
// 		err := rows.Scan(&voter.ID, &voter.PoliticianID, &voter.FirstName, &voter.LastName, &voter.Gender, &voter.Age)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		result = append(result, voter)
// 	}

// 	json.NewEncoder(w).Encode(result)
// }
