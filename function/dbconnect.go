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

type Votings struct {
	VoterId      int
	PoliticianId int
	FirstName    string
	LastName     string
	Gender       string
	Age          int
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/politicians_voting")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("sukses konek database")
	// fmt.Println(db)
	return db, nil
}

func GetPoliticians() (dataPoliticians []Politicians) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var dataPoliticians []Politicians

	data, err := db.Query("SELECT * FROM politicians")

	for data.Next() {
		var dat Politicians
		if data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		dataPoliticians = append(dataPoliticians, dat)
	}
	return
}

func GetVotings() (dataVotings []Votings) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var dataVotings []Votings

	data, err := db.Query("SELECT * FROM voting")

	for data.Next() {
		var dat Votings
		if data.Scan(&dat.VoterId, &dat.PoliticianId, &dat.FirstName, &dat.LastName, &dat.Gender, &dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}
		dataVotings = append(dataVotings, dat)
	}
	return
}
