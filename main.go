package main

import (
	"fmt"
	"net/http"
	"webservice/routing"
)

func main() {
	http.HandleFunc("/voting", routing.ShowVoting)
	http.HandleFunc("/votingbymale", routing.ShowVotingbyMale)
	http.HandleFunc("/votingbyfemale", routing.ShowVotingbyFemale)
	http.HandleFunc("/politician", routing.ShowPolitician)

	port := "localhost:3000"

	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
