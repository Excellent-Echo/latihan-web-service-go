package local

import (
	"LATIHAN-WEB-SERVICE-GO/connect"
	"LATIHAN-WEB-SERVICE-GO/entiti"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func decodePolitician(file string) entiti.Politician {
	var politik entiti.Politician

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &politik)
	if err != nil {
		fmt.Println(err.Error())
	}
	return politik

}

func masukinPolitician(p entiti.Politician) {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, nile := range p {
		_, err = db.Exec("INSERT INTO politicians (name,party,location,grade_current) VALUES (?,?,?,?)", nile.Nama, nile.Party, nile.Lokasi, nile.Grade)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("insert data sukses banget")

}
func Eksekusi() {
	satpol := decodePolitician("dataJson/politicians.json")
	masukinPolitician(satpol)
}
