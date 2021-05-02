package webService

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

func webService() {
	var uri = ""https://www.google.com/search?q=golang&rlz=1C1GCEA_enID937ID937&oq=golang&aqs=chrome..69i57j69i59l3j46j0l5.1322j0j15&sourceid=chrome&ie=UTF-8"

	url, err := url.Parse(uri)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(url.Query())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		t, err :=template.ParseFiles("temleate.html")
	})
}