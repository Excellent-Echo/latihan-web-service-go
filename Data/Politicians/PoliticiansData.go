package Politicians

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Potilicians []struct {
	PoliticianId 	int 	`json:"politician_id"`
	Name 			string 	`json:"name"`
	Party 			string 	`json:"party"`
	Location 		string 	`json:"location"`
	GradeCurrent 	float32	`json:"grade_current"`
}

func PoliticianData() {
	var politiciansData Potilicians
	jsonFile, err := os.Open("Data/Politicians/politicians.json")
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file politicians.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politiciansData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range politiciansData {
		fmt.Println(data.PoliticianId)
		fmt.Println(data.Name)
		fmt.Println(data.Party)
		fmt.Println(data.Location)
		fmt.Println(data.GradeCurrent)
	}
}