package main

import (
	"latihan-web-service-go/routing"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// helper.Candidate()
	// helper.Voter()
	routing.Routing()
}
