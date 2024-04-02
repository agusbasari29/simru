package seeders

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/sip/simru/entity"
	"golang.org/x/crypto/bcrypt"
)

type UsersSeeder struct {
	Name           string `faker:"name"`
	Username       string `faker:"username"`
	Email          string `faker:"email"`
	Phone          string `faker:"phone_number"`
	CompanyName    string `faker:"word"`
	CompanyAddress string `faker:"sentence"`
}

func UsersSeederUp() {
	seeder := UsersSeeder{}
	users := entity.Users{}
	person := entity.Persons{}
	number := 18
	for i := 0; i < number; i++ {
		err := faker.FakeData(&seeder)
		if err != nil {
			fmt.Printf("%+v", err)
		}
		person.Name = seeder.Name
		j := rand.Intn(9999999)
		nip := strconv.Itoa(j)
		person.NIP = nip
		person.Email = seeder.Email
		person.Phone = seeder.Phone
		person.CompanyName = seeder.CompanyName
		person.CompanyAddress = seeder.CompanyAddress
		person.CreatedAt = time.Now()
		res, err := personRepo.CreatePerson(person)
		if err != nil {
			fmt.Println("Error when create for new person")
		}
		password, _ := bcrypt.GenerateFromPassword([]byte("sangatrahasia"), bcrypt.DefaultCost)
		users.Username = seeder.Username
		users.Password = string(password)
		roleId := uint64(i + 1)
		users.UserRoleID = roleId
		users.PersonID = res.ID
		users.CreatedAt = time.Now()
		userRepo.CreateUser(users)
	}
}
