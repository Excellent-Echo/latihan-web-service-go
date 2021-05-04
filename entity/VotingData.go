package entity
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"os"
//)
//
//type Voters []struct {
//	VoterId 		int 	`json:"voter_id"`
//	PoliticianId 	int 	`json:"policitian_id"`
//	FirstName 		string 	`json:"first_name"`
//	LastName 		string 	`json:"last_name"`
//	Gender 			string 	`json:"gender"`
//	Age				int 	`json:"age"`
//}
//
//func VoterData() (Voters, error) {
//	var votersData Voters
//
//	jsonFile, err := os.Open("data/voting.json")
//	if err != nil{
//		return Voters{}, err
//	}
//	fmt.Println("success get file voting.json")
//	defer jsonFile.Close()
//
//	byteData, _ := ioutil.ReadAll(jsonFile)
//	err = json.Unmarshal(byteData, &votersData)
//	if err != nil{
//		return Voters{}, err
//	}
//	return votersData, nil
//
//	for _, data := range politiciansData {
//		fmt.Println(data.VoterId)
//		fmt.Println(data.PoliticianId)
//		fmt.Println(data.FirstName)
//		fmt.Println(data.LastName)
//		fmt.Println(data.Gender)
//		fmt.Println(data.Age)
//	}
//}