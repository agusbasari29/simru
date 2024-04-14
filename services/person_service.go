package services

import (
	"time"

	"github.com/mashingan/smapping"
	"github.com/sip/simru/entity"
	"github.com/sip/simru/repository"
	"github.com/sip/simru/request"
)

type PersonService interface {
	CreatePerson(request request.UserRegisterRequest) (entity.Persons, error)
	GetPerson(request request.PersonRequest) (entity.Persons, error)
}

type personService struct {
	personRepository repository.PersonRepository
}

func NewPersonService(personRepository repository.PersonRepository) *personService {
	return &personService{personRepository}
}

func (s *personService) CreatePerson(request request.UserRegisterRequest) (entity.Persons, error) {
	person := entity.Persons{}
	err := smapping.FillStruct(&person, smapping.MapFields(&request))
	smapError(err)
	person.CreatedAt = time.Now()
	newPerson, err := s.personRepository.CreatePerson(person)
	if err != nil {
		return newPerson, err
	}
	return newPerson, nil
}

func (s *personService) GetPerson(request request.PersonRequest) (entity.Persons, error) {
	person := entity.Persons{}
	err := smapping.FillStruct(&person, smapping.MapFields(&request))
	smapError(err)
	result, err := s.personRepository.GetPerson(person)
	if err != nil {
		return result, err
	}
	return result, nil
}
