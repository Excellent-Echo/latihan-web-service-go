// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// )

// type VotingData struct {
// 	Voters []Data `json:"pemilih"`
// }
// type Data struct {
// 	Votersid     int    `json:"voters_id"`
// 	PoliticanId  int    `json:"policitian_id"`
// 	NamaDepan    string `json:"first_name"`
// 	NamaBelakang string `json:"last_name"`
// 	JenisKelamin string `json:"gender"`
// 	Umur         int    `json:"age"`
// }

// func main() {
// 	jsonFile, err := os.Open("dataJson/voting.json")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("succes open jsonFile")
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)

// 	var voters VotingData

// 	json.Unmarshal([]byte(byteValue), &voters)

// 	for i := 0; i < len(voters.Voters); i++ {
// 		fmt.Println("voters Votersid: " + strconv.Itoa(voters.Voters[i].Votersid))
// 		fmt.Println("voters PoliticanId: " + strconv.Itoa(voters.Voters[i].PoliticanId))
// 		fmt.Println("voters NamaDepan: " + voters.Voters[i].NamaDepan)
// 		fmt.Println("voters NamaBelakang: " + voters.Voters[i].NamaBelakang)
// 		fmt.Println("voters Jeniskelamin: " + voters.Voters[i].JenisKelamin)
// 		fmt.Println("voters Umur: " + strconv.Itoa(voters.Voters[i].Umur))
// 	}

// }
