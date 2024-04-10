package services

import (
	"log"
	"time"

	"github.com/mashingan/smapping"
	"github.com/sip/simru/entity"
	"github.com/sip/simru/repository"
	"github.com/sip/simru/request"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	IsExist(username entity.Users) bool
	RegisterUser(request request.UserRegisterRequest) (entity.Users, error)
	VerifyCredential(request request.AuthLoginRequest) interface{}
	GetUser(request request.UserRequest) (entity.Users, error)
}

type userService struct {
	userRepository   repository.UserRepository
	personRepository repository.PersonRepository
}

func NewUserService(userRepository repository.UserRepository, personRepository repository.PersonRepository) *userService {
	return &userService{userRepository, personRepository}
}

func (s *userService) IsExist(username entity.Users) bool {
	_, err := s.userRepository.GetUser(username)
	return err != nil
}

func (s *userService) RegisterUser(request request.UserRegisterRequest) (entity.Users, error) {
	user := entity.Users{}
	person := entity.Persons{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	smapError(err)
	err = smapping.FillStruct(&person, smapping.MapFields(&request))
	smapError(err)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	newPerson, err := s.personRepository.CreatePerson(person)
	if err != nil {
		log.Fatalf("%v", err)
	}
	user.PersonID = newPerson.ID
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) VerifyCredential(request request.AuthLoginRequest) interface{} {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	smapError(err)
	result, err := s.userRepository.GetUser(user)
	if err != nil {
		return false
	} else {
		comparedPassword := comparePassword(result.Password, user.Password)
		if result.Username == user.Username && comparedPassword {
			return result
		}
	}
	return false
}

func (s *userService) GetUser(request request.UserRequest) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	smapError(err)
	result, err := s.userRepository.GetUser(user)
	if err != nil {
		return result, err
	}
	return result, nil
}

func comparePassword(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

func smapError(err error) {
	if err != nil {
		log.Fatalf("Failed to map %v.", err)
	}
}
