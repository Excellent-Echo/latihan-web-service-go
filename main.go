package main

import (
	"latihan-web-service-go/webserver"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// helper.Candidate()
	// helper.Voter()
	// routing.Routing()
	webserver.WebService()
}
