package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type User struct {
	Politik []Politician `json:"politik"`
}

type Politician struct {
	Politician_id int    `json:"politician_id"`
	Name          string `json:"name"`
	Party         string `json:"party"`
	Location      string `json:"location"`
	Grade_current int    `json:"grade_current"`
}

func main() {
	jsonFilePolitician, err := os.Open("politicians.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success open jsonFilePolitician")

	defer jsonFilePolitician.Close()

	byteValue, _ := ioutil.ReadAll(jsonFilePolitician)

	var userpolitik User

	json.Unmarshal([]byte(byteValue), &userpolitik)

	for i := 0; i < len(userpolitik.Politik); i++ {
		fmt.Println("politian_id: " + strconv.Itoa(userpolitik.Politik[i].Politician_id))
		fmt.Println("name: " + userpolitik.Politik[i].Name)
		fmt.Println("party: " + userpolitik.Politik[i].Party)
		fmt.Println("location: " + userpolitik.Politik[i].Location)
		fmt.Println("grade_current: " + strconv.Itoa(userpolitik.Politik[i].Grade_current))
	}
}
