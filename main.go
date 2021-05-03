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
	// webserver.WebServer()
	http.HandleFunc("/", webserver.RootHandler)
	http.HandleFunc("/votings", webserver.AllVotingHandler)

	port := "localhost:8080"
	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
