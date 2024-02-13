package dto

type People struct {
	ID         int
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	Country    []Country
}

type UpdatePeople struct {
	ID         int
	Name       string
	Surname    string
	Patronymic string
}

type DeletePeople struct {
	ID int
}

type CreatePeople struct {
	Name       string
	Surname    string
	Patronymic string
}

type Country struct {
	ID          int
	CountryID   string
	Probability float64
}
