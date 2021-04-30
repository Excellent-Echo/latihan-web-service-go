package function

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Politicians struct {
	PoliticianId int
	Name         string
	Party        string
	Location     string
	GradeCurrent float32
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/politicians_voting")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("sukses konek database")
	return db, nil
}

func getPoliticians() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dataPoliticians []Politicians

	data, err := db.Query("SELECT * FROM politicians")

	for data.Next() {
		var dat Politicians
		if data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		dataPoliticians = append(dataPoliticians, dat)
	}
}
