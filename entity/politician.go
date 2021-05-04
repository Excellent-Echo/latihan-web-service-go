package entity

type Politician struct {
	PoliticianID   int     `json:"politician_id"`
	NamePolitician string  `json:"name"`
	Party          string  `json:"party"`
	Location       string  `json:"location"`
	GradeCurrent   float32 `json:"grade_current"`
	ResVote        float32
}

type Politicians []Politician
