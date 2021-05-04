package function

import "fmt"

type PoliticiansWithVoting struct {
	Politicians
	Voting int
}

func QueryPoliticiansWithVoting() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryVotingMale() (MaleVotings []Votings) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voting WHERE gender='male'")

	// var datas []Votings

	for data.Next() {
		var dat Votings
		if data.Scan(&dat.VoterId, &dat.PoliticianId, &dat.FirstName, &dat.LastName, &dat.Gender, &dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}
		MaleVotings = append(MaleVotings, dat)
	}
	return
}

func QueryVotingFemale() (FemaleVotings []Votings) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voting WHERE gender='female'")

	// var datas []Votings

	for data.Next() {
		var dat Votings
		if data.Scan(&dat.VoterId, &dat.PoliticianId, &dat.FirstName, &dat.LastName, &dat.Gender, &dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}
		FemaleVotings = append(FemaleVotings, dat)
	}
	return
}

func QueryVotingMaleAndNameB() {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM voting WHERE gender='male' AND first_name LIKE 'B%';")

	var datas []Votings

	for data.Next() {
		var dat Votings
		if data.Scan(&dat.VoterId, &dat.PoliticianId, &dat.FirstName, &dat.LastName, &dat.Gender, &dat.Age); err != nil {
			fmt.Println(err.Error())
			return
		}
		datas = append(datas, dat)
	}
}

func QueryPoliticians1stVotingLocIL() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='IL' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingLocNY() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='NY' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingLocTX() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='TX' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingLocLA() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='LA' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingLocWA() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='WA' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingLocFL() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='FL' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingLocHI() (datas []PoliticiansWithVoting) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	data, err := db.Query("SELECT p.politician_id, p.name, p.party, p.location, p.grade_current, COUNT(v.politician_id) AS voting FROM politicians AS p JOIN voting AS v ON p.politician_id = v.politician_id WHERE location='HI' GROUP BY p.politician_id;")

	// var datas []PoliticiansWithVoting

	for data.Next() {
		var dat = PoliticiansWithVoting{}
		err := data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent, &dat.Voting)
		if err != nil {
			fmt.Println(err.Error())
		}
		datas = append(datas, dat)
	}
	return
}

func QueryPoliticians1stVotingNY() {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	data, err := db.Query("SELECT * FROM politicians WHERE location='NY' ORDER BY grade_current DESC LIMIT 1;")

	var datas []Politicians

	for data.Next() {
		var dat Politicians
		if data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		datas = append(datas, dat)
	}
}

func QueryPoliticians3rdVoting() {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}

	data, err := db.Query("SELECT * FROM politicians ORDER BY grade_current DESC LIMIT 3;")

	var datas []Politicians
	for data.Next() {
		var dat Politicians
		if data.Scan(&dat.PoliticianId, &dat.Name, &dat.Party, &dat.Location, &dat.GradeCurrent); err != nil {
			fmt.Println(err.Error())
			return
		}
		datas = append(datas, dat)
	}
}
