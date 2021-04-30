package Data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//type Politicians struct {
//	Politicians []Politician `json:"politicians"`
//}

type Potilician []struct {
	Politician_Id 	int`json:"politician"`
	Name 			string `json:"name"`
	Party 			string `json:"party"`
	Location 		string `json:"location"`
	Grade_Current 	int `json:"grade_current"`
}

type Voting struct {
	Voter_Id 	int `json:"voter_id"`
	Potilician 	Potilician `json:"potilician"`
	First_Name 	string `json:"first_name"`
	Last_Name 	string `json:"last_name"`
	Gender 		string `json:"gender"`
	Age 		int `json:"age"`
}

func jsonData() {
	var politiciansData Potilician
	jsonFile, err := os.Open("politicians.json")
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Println("succes get file politicians.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(politiciansData)
	err = json.Unmarshal(byteData, &politiciansData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range politiciansData.Politician {
		fmt.Println("politician_id :", data.Politician_Id)
		fmt.Println("name :", data.Name
		fmt.Println("party :", data.Party)
		fmt.Println("location :", data.Location)
		fmt.Println("grade_current :", data.Grade_Current)
		)
	}
}