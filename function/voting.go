package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Voting []struct {
	VoterId      int    `json:"voter_id"`
	PoliticianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func ModelVoting() {
	jsonVoting, err := os.Open("voting.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer jsonVoting.Close()

	byteVoting, _ := ioutil.ReadAll(jsonVoting)

	var voting Voting
	json.Unmarshal([]byte(byteVoting), &voting)
	// fmt.Println("Data Voting JSON")
	// for _, val := range voting {
	// 	fmt.Println(val)
	// }

	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//insert data Voting to database
	for _, val := range voting {
		_, err = db.Exec("insert into voting values(?,?,?,?,?,?)", val.VoterId, val.PoliticianId, val.FirstName, val.LastName, val.Gender, val.Age)
		fmt.Println("sukses post data voting ke database")
	}

}
