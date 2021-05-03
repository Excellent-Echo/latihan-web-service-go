package main

import (
	"fmt"
	"latihan-web-service-go/webserver"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// helper.Candidate()
	// helper.Voter()
	// routing.Routing()
	http.HandleFunc("/", webserver.RootHandler)
	http.HandleFunc("/votings", webserver.AllVotingHandler)
	http.HandleFunc("/votings_male", webserver.MaleVotingHandler)
	http.HandleFunc("/votings_female", webserver.FemaleVotingHandler)
	http.HandleFunc("/politicians", webserver.AllPoliticianHandler)
	http.HandleFunc("/politicians_voting", webserver.AllPoliticianWithVoteHandler)
	http.HandleFunc("/politicians_il", webserver.ILPoliticiansHandler)
	http.HandleFunc("/politicians_ny", webserver.NYPoliticiansHandler)
	http.HandleFunc("/politicians_tx", webserver.TXPoliticiansHandler)
	http.HandleFunc("/politicians_la", webserver.LAPoliticiansHandler)
	http.HandleFunc("/politicians_wa", webserver.WAPoliticiansHandler)
	http.HandleFunc("/politicians_fl", webserver.FLPoliticiansHandler)
	http.HandleFunc("/politicians_hi", webserver.HIPoliticiansHandler)

	port := "localhost:8080"
	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
