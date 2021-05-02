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
	Politik []Data `json`
}
type Data struct {
	Politian_id   string `json : "politician_id"`
	Name          string `json : "name"`
	Party         string `json : "party"`
	Location      string `json : "location"`
	Grade_current int    `json : "grade_current"`
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

	var Politik PolitianData

	json.Unmarshal(byteValue, &Politik)
	fmt.Println("Politik location: " + Politik.Politik[0].Location)

	for i := 0; i < len(Politik.Politik); i++ {
		fmt.Println("Politik politian_id: " + Politik.Politik[i].Politian_id)
		fmt.Println("Politik name: " + Politik.Politik[i].Name)
		fmt.Println("Politik party: " + Politik.Politik[i].Party)

		fmt.Println("Politik grade_current: " + strconv.Itoa(Politik.Politik[i].Grade_current))

	}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application")
	// })
}
