package main

import (
	"fmt"
	"latihan-web-service-go/handler"
	"net/http"
)

func main() {
	//helper.Politicians()
	//helper.Voter()
	//router.Endpoint()

	// Semua endpoint
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/votings", handler.VotingHandler)
	http.HandleFunc("/votings_male", handler.VotingMaleHandler)
	http.HandleFunc("/votings_female", handler.VotingFemaleHandler)
	http.HandleFunc("/politicians", handler.PoliticHandler)
	http.HandleFunc("/politicians_voting", handler.PoliticAndVoting)
	http.HandleFunc("/politicians_il", handler.PoliticILLocation)
	http.HandleFunc("/politicians_ny", handler.PoliticNYLocation)
	http.HandleFunc("/politicians_tx", handler.PoliticTXLocation)
	http.HandleFunc("/politicians_la", handler.PoliticLALocation)
	http.HandleFunc("/politicians_wa", handler.PoliticWALocation)
	http.HandleFunc("/politicians_fl", handler.PoliticFLLocation)
	http.HandleFunc("/politicians_hi", handler.PoliticHILocation)

	// untuk menyambungkan localhost dari local komputer
	port := "localhost:3000"
	fmt.Println("server running on port", port)
	http.ListenAndServe(port, nil)
}
