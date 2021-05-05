package local

import (
	"LATIHAN-WEB-SERVICE-GO/connect"
	"LATIHAN-WEB-SERVICE-GO/entiti"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func decodeVoting(poting string) entiti.Votings {
	var vote entiti.Votings

	jsonFile, err := os.Open(poting)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &vote)
	if err != nil {
		fmt.Println(err.Error())
	}
	return vote

}

func MasukinVote(Vot entiti.Votings) {
	db, err := connect.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	for _, nileVot := range Vot {
		_, err = db.Exec("INSERT INTO Voting (politician_id,first_name,last_name,gender,age) VALUES (?,?,?,?,?)", nileVot.PoliticanId, nileVot.NamaDepan, nileVot.NamaBelakang, nileVot.JenisKelamin, nileVot.Umur)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("insert data sukses bet")

}
func ExecuteVot() {
	pot := decodeVoting("dataJson/voting.json")
	MasukinVote(pot)
}
