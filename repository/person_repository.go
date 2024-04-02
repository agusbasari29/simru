package repository

import "gorm.io/gorm"

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepository {
	return &personRepository{db}
}
