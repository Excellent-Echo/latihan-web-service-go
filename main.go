package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//politian_id, name, party, location, grade_current
// //   "politician_id" : 1,
// "name" : "Aaron Schock",
// "party" : "REPUBLIC",
// "location" : "IL",
// "grade_current"
type PolitianData struct {
	Politik []Data `json:"politik"`
}
type Data struct {
	Politian_id   int    `json:"politician_id"`
	Name          string `json:"name"`
	Party         string `json:"party"`
	Location      string `json:"location"`
	Grade_current int    `json:"grade_current"`
}

func main() {
	jsonFile, err := os.Open("dataJson/politicians.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("succes open jsonFile")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var politik PolitianData

	json.Unmarshal([]byte(byteValue), &politik)
	// fmt.Println(byteValue)

	for i := 0; i < len(politik.Politik); i++ {
		fmt.Println("Politik politian_id: " + strconv.Itoa(politik.Politik[i].Politian_id))
		fmt.Println("Politik name: " + politik.Politik[i].Name)
		fmt.Println("Politik party: " + politik.Politik[i].Party)
		fmt.Println("Politik location: " + politik.Politik[i].Location)
		fmt.Println("Politik grade_current: " + strconv.Itoa(politik.Politik[i].Grade_current))

	}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application")
	// })
}
