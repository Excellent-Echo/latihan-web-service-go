package tables_data

type Politician struct {
	PoliticianID int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}