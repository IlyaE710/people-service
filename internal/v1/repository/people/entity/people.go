package entity

import "gorm.io/gorm"

type People struct {
	gorm.Model
	ID         int
	Surname    string
	Name       string
	Patronymic string
	Age        int
	Gender     string
	Country    []Country `gorm:"foreignKey:PeopleID"`
}

type Country struct {
	gorm.Model
	ID          int
	PeopleID    int
	CountryID   string
	Probability float64
}
