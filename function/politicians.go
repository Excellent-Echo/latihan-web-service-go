package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Politician []struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

func ModelPoliticians() {
	//Politicians JSON
	jsonPoliticians, err := os.Open("politicians.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println("success open jsonPoliticians")

	defer jsonPoliticians.Close()

	bytePocisians, _ := ioutil.ReadAll(jsonPoliticians)

	var politician Politician

	json.Unmarshal([]byte(bytePocisians), &politician)

	// fmt.Println("Data Politician JSON")
	// for _, val := range politician {
	// 	fmt.Println(val)
	// }
	//connect database
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//insert data politicians to databse
	for _, val := range politician {
		_, err = db.Exec("INSERT INTO politicians VALUES(?, ?, ?, ?, ?)", val.PoliticianId, val.Name, val.Party, val.Location, val.GradeCurrent)
		fmt.Println("sukses post ke database")
	}
}
