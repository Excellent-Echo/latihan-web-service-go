package helper

import (
	"fmt"
	"html/template"
	"net/http"
)

func WebService() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"nama":  "Muhamad Aziz",
			"pesan": "Selamat datang di web service go",
		}

		t, err := template.ParseFiles("helper/template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})

	port := "localhost:5000"

	fmt.Println("Starting at server port", port)

	http.ListenAndServe(port, nil)
}
