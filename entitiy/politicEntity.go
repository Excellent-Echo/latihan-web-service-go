package entitiy

// PoliticVoterDataJoin data struct
type PoliticVoterDataJoin struct {
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
	VoterId      int     `json:"voter_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Gender       string  `json:"gender"`
	Age          int     `json:"age"`
}
type PoliticVotersJoin []PoliticVoterDataJoin

// Politic data struct
type Politic struct {
	PoliticianId int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}
type Politician []Politic
