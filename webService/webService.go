package WebService

import (
	"fmt"
	"latihan-web-service-go/function"
	"net/http"
	"text/template"
)

func Service() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/index.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, nil)
	})

	http.HandleFunc("/votings", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/votings.html")

		// data := map[string]interface{}
		var datas = function.GetVotings()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		firstname := datas[0].FirstName
		header := "<h2>" + firstname + "</h2>"

		t.Execute(w, header)
	})

	http.HandleFunc("/votings_male", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/votings.html")

		data := map[string]string{
			"header":  "<h2>Voting Male asdasdasdasd Pages</h2>",
			"name":    "male",
			"message": "asique",
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)

	})

	http.HandleFunc("/votings_female", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/votings.html")

		data := map[string]string{
			"header":  "votings female",
			"name":    "female",
			"message": "asique",
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)

	})

	http.HandleFunc("/politicians", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_voting", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians voting pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_il", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians il pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_ny", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians ny pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_tx", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians tx pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_la", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians la pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_wa", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians wa pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_fl", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians fl pages",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/politicians_hi", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("webService/pages/politicians_loc.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"header": "politicians hi pages",
		}
		t.Execute(w, data)
	})

	port := "localhost:8080"
	fmt.Println("Starting Web Service at", port)
	http.ListenAndServe(port, nil)
}
