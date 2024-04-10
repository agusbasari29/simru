package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type PersonRepository interface {
	CreatePerson(person entity.Persons) (entity.Persons, error)
	UpdatePerson(person entity.Persons) (entity.Persons, error)
	GetPerson(person entity.Persons) (entity.Persons, error)
	GetPersons() ([]entity.Persons, error)
	DeletePerson(person entity.Persons) bool
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepository {
	return &personRepository{db}
}

func (r *personRepository) CreatePerson(person entity.Persons) (entity.Persons, error) {
	err := r.db.Create(&person).Error
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *personRepository) UpdatePerson(person entity.Persons) (entity.Persons, error) {
	err := r.db.Model(&entity.Persons{ID: person.ID}).Save(&person).Error
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *personRepository) GetPerson(person entity.Persons) (entity.Persons, error) {
	err := r.db.First(&person).Error
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *personRepository) GetPersons() ([]entity.Persons, error) {
	persons := []entity.Persons{}
	err := r.db.Find(&persons).Error
	if err != nil {
		return persons, err
	}
	return persons, nil
}

func (r *personRepository) DeletePerson(person entity.Persons) bool {
	return r.db.Delete(&person).Error == nil
}
