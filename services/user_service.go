package services

import (
	"fmt"
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
	RegisterUser(request request.UserRegisterRequest, personID uint64) (entity.Users, error)
	VerifyCredential(request request.AuthLoginRequest) interface{}
	GetUser(request request.UserRequest) (entity.Users, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) IsExist(username entity.Users) bool {
	_, err := s.userRepository.GetUser(username)
	return err != nil
}

func (s *userService) RegisterUser(request request.UserRegisterRequest, personID uint64) (entity.Users, error) {
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
	user.PersonID = personID
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
	fmt.Println(request)
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
