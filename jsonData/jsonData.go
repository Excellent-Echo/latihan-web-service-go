package jsonData

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Politicians []struct {
	Politicians_id int     `json:"politician_id"`
	Name           string  `json:"name"`
	Party          string  `json:party`
	Location       string  `json:location`
	Grade_Current  float32 `json:grade_current`
}

func jsonData() {
	var politiciansData Politicians
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file data.json")
	defer jsonFile.Close()
}

byteData,_ := ioutil.ReadAll(jsonFile)
err = json.Unmarshal (byteData, &politiciansData)

if err != nill {
	fmt.Println(err.Error())
	return
}

for _,data := range politiciansData.Politicians {
	fmt.Println("politician_id", data.Politicians_id)
	fmt.Println("name :", data.Name)
	fmt.Println("party :", data.Party)
	fmt.Println("location :", data.Location)
	fmt.Println("grade_current", data.Grade_Current)
}
