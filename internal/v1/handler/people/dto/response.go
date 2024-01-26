package dto

type ErrorResponse struct {
	Message string `json:"message"`
}

type PeopleResponse struct {
	ID         int               `json:"id"`
	Name       string            `json:"name"`
	Surname    string            `json:"surname"`
	Patronymic string            `json:"patronymic"`
	Age        int               `json:"age"`
	Gender     string            `json:"gender"`
	Country    []CountryResponse `json:"country"`
}

type CountryResponse struct {
	ID          int     `json:"id"`
	CountryID   string  `json:"country-id"`
	Probability float64 `json:"probability"`
}
