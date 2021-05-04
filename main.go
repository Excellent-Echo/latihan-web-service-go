package main

import (
	"fmt"
	"net/http"
	"webservice/routing"
)

func main() {
	http.HandleFunc("/voting", routing.ShowVoting)
	http.HandleFunc("/votingbymale", routing.ShowVotingbyMale)

	port := "localhost:3000"

	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
