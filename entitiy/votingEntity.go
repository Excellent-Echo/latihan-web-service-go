package entitiy

type VoterData struct {
	VoterId      int    `json:"voter_id"`
	PolicitianId int    `json:"policitian_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}
type Voters []VoterData
