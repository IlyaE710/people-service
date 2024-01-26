package reponse

type People struct {
	Name    string        `json:"name,omitempty"`
	Age     int           `json:"age,omitempty"`
	Gender  string        `json:"gender,omitempty"`
	Country []CountryInfo `json:"country,omitempty"`
}

type CountryInfo struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
