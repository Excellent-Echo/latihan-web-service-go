package Data

import "fmt"

type VotersDB struct {
	Voter_id      int
	Politician_id int
	First_name    string
	Last_name     string
	Gender        string
	Age           int
}

//get data Politicians by ID
func GetDataVoterById(id int) (dataVoterDB VotersDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voters WHERE voter_id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	for data.Next() {

		var dat VotersDB

		if data.Scan(
			&dat.Voter_id,
			&dat.Politician_id,
			&dat.First_name,
			&dat.Last_name,
			&dat.Gender,
			&dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataVoterDB = dat
	}
	return
}

//Get All data Politicians
func GetDataVoters() (dataVotersDB []VotersDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voters")
	if err != nil {
		fmt.Println(err.Error())
	}

	for data.Next() {

		var dat VotersDB

		if data.Scan(
			&dat.Voter_id,
			&dat.Politician_id,
			&dat.First_name,
			&dat.Last_name,
			&dat.Gender,
			&dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}

		dataVotersDB = append(dataVotersDB, dat)
	}
	return
}

//this function for initial post from json only do not use for others
func InitialPostDataVoters(id int, id_politician int, first_name string, last_name string, gender string, age int) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//query
	_, err = db.Exec("INSERT INTO voters VALUES (?,?,?,?,?,?)", id, id_politician, first_name, last_name, gender, age)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Success initial insert data")
}
