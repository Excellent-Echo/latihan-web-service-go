package Data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// POLITICIANS DATA

type Politicians []struct {
	Politician_id int     `json:"politician_id"`
	Name          string  `json:"name"`
	Party         string  `json:"party"`
	Location      string  `json:"location"`
	Grade_current float32 `json:"grade_current"`
}

func PoliticianData() {

	var politiciansData Politicians

	//get data json
	jsonFile, err := os.Open("Data/politicians.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file poticians.json")
	defer jsonFile.Close()

	//read data json
	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &politiciansData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// for _, data := range politiciansData {
	// 	fmt.Println(data.Politician_id)
	// 	fmt.Println(data.Name)
	// 	fmt.Println(data.Party)
	// 	fmt.Println(data.Location)
	// 	fmt.Println(data.Grade_current)
	// }
}
