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

//Get All data Voters
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

//Get All data Voters with Bname
func GetDataMaleVotersWithBname() (dataVotersDB []VotersDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voters where gender='male' AND first_name LIKE 'B%'")
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

//Get All Male Voters
func GetDataMaleVoters() (dataVotersDB []VotersDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voters where gender='male'")
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

//Get All Female Voters
func GetDataFemaleVoters() (dataVotersDB []VotersDB) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voters where gender='female'")
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
