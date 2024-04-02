package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepository {
	return &personRepository{db}
}

func (r *personRepository) CreatePerson(person entity.Persons) (entity.Persons, error) {
	err := r.db.Raw("INSERT INTO persons (name, nip, company_name, company_address, email, phone, created_at) VALUE (@Name, @NIP, @CompanyName, @CompanyAddress, @Email, @Phone, @CreatedAt)", person).Create(&person).Error
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *personRepository) GetPerson() {

}
