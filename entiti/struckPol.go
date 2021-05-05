package entiti

type Politicians struct {
	PoliticianId int     `json:"politician_id"`
	Nama         string  `json:"name"`
	Party        string  `json:"party"`
	Lokasi       string  `json:"location"`
	Grade        float32 `json:"grade_current"`
}

type Politician []Politicians
