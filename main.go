package main

import (
	"latihan-web-service-go/routing"

	_ "github.com/go-sql-driver/mysql"
)

// func politicianRoute(p Politicians) {
// 	http.HandleFunc("/politicians", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		if r.Method == "GET" {
// 			byteJson, err := json.Marshal(p)

// 			if err != nil {
// 				http.Error(w, "error internal server", http.StatusInternalServerError)
// 				return
// 			}

// 			w.Write(byteJson)
// 			return
// 		}

// 		http.Error(w, "", http.StatusBadRequest)
// 	})

// port := "localhost:4444"

// 	fmt.Println("server running on port", port)

// 	http.ListenAndServe(port, nil)
// }

func main() {
	// helper.Candidate()
	// helper.Voter()
	routing.Routing()
	// politicianRoute(p)
}
