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
	http.HandleFunc("/showpoliticianjoinvoting", routing.ShowPoliticianjoinVoting)
	http.HandleFunc("/showpoliticianjoinvotingbyil", routing.ShowPoliticianjoinVotingbyLocationIL)
	http.HandleFunc("/showpoliticianjoinvotingbyny", routing.ShowPoliticianjoinVotingbyLocationNY)
	http.HandleFunc("/showpoliticianjoinvotingbytx", routing.ShowPoliticianjoinVotingbyLocationTX)
	http.HandleFunc("/showpoliticianjoinvotingbyla", routing.ShowPoliticianjoinVotingbyLocationLA)
	http.HandleFunc("/showpoliticianjoinvotingbywa", routing.ShowPoliticianjoinVotingbyLocationWA)
	http.HandleFunc("/showpoliticianjoinvotingbyfl", routing.ShowPoliticianjoinVotingbyLocationFL)
	http.HandleFunc("/showpoliticianjoinvotingbyhi", routing.ShowPoliticianjoinVotingbyLocationHI)

	port := "localhost:3000"

	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
