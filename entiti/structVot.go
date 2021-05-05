package entiti

type Voting struct {
	Votersid     int    `json:"voter_id"`
	PoliticanId  int    `json:"policitian_id"`
	NamaDepan    string `json:"first_name"`
	NamaBelakang string `json:"last_name"`
	JenisKelamin string `json:"gender"`
	Umur         int    `json:"age"`
}

type VotingS []Voting
