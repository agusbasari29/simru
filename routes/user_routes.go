package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sip/simru/database"
	"github.com/sip/simru/handler"
	"github.com/sip/simru/helper"
	"github.com/sip/simru/repository"
	"github.com/sip/simru/services"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	db := database.SetupDBConnection()
	userRepository := repository.NewUserRepository(db)
	personRepository := repository.NewPersonRepository(db)
	userRolesRepository := repository.NewUserRoleRepository(db)
	sectionsRepository := repository.NewSectionRepository(db)
	userService := services.NewUserService(userRepository)
	jwtService := services.NewJWTService()
	personService := services.NewPersonService(personRepository)
	userRolesServices := services.NewUserRolesServices(userRolesRepository)
	sectionsServices := services.NewSectionsServices(sectionsRepository)
	userHandler := handler.NewUserHandler(userService, jwtService, personService, userRolesServices, sectionsServices)
	return []helper.Route{
		{
			Method:  "GET",
			Path:    "/user",
			Handler: []gin.HandlerFunc{userHandler.GetUser},
		},
	}
}
