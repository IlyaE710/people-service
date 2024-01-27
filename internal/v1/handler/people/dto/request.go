package dto

type CreatePeopleRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UpdatePeopleRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
}

type CountryRequest struct {
	ID          int     `json:"id"`
	CountryID   string  `json:"country-id"`
	Probability float64 `json:"probability"`
}
