package seeders

import (
	"github.com/sip/simru/database"
	"github.com/sip/simru/repository"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB = database.SetupDBConnection()
	sectionRepo           = repository.NewSectionRepository(db)
	userRoleRepo          = repository.NewUserRoleRepository(db)
	personRepo            = repository.NewPersonRepository(db)
	userRepo              = repository.NewUserRepository(db)
)

func Seeders() {
	SectionSeedersUp()
}
