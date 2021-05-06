package main

import (
	"electionGo/query"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//SEEDING DATA (RELEASE 1)
	// data.DecodePoliticians()
	// data.DecodeVoters()

	//QUERY DATA (RELEASE 2)
	// query.Release2_1()
	// query.Release2_2()
	// query.Release2_3()
	// query.Release2_4()

	//WEB SERVICE (RELEASE 3)
	router := mux.NewRouter()

	router.HandleFunc("/voters", query.GetAllVoters).Methods("GET")
	router.HandleFunc("/voters_male", query.GetAllMaleVoters).Methods("GET")
	router.HandleFunc("/voters_female", query.GetAllFemaleVoters).Methods("GET")
	router.HandleFunc("/politicians", query.GetAllPoliticians).Methods("GET")
	router.HandleFunc("/politicians_voting", query.GetAllPoliticiansVoting).Methods("GET")
	router.HandleFunc("/politicians_il", query.GetAllPoliticiansIL).Methods("GET")
	router.HandleFunc("/politicians_ny", query.GetAllPoliticiansNY).Methods("GET")
	router.HandleFunc("/politicians_tx", query.GetAllPoliticiansTX).Methods("GET")
	router.HandleFunc("/politicians_la", query.GetAllPoliticiansLA).Methods("GET")
	router.HandleFunc("/politicians_wa", query.GetAllPoliticiansWA).Methods("GET")
	router.HandleFunc("/politicians_fl", query.GetAllPoliticiansFL).Methods("GET")
	router.HandleFunc("/politicians_hi", query.GetAllPoliticiansHI).Methods("GET")

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", router)

}
