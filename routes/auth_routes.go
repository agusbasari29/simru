package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sip/simru/database"
	"github.com/sip/simru/handler"
	"github.com/sip/simru/helper"
	"github.com/sip/simru/repository"
	"github.com/sip/simru/services"
)

type AuthRoutes struct{}

func (r AuthRoutes) Route() []helper.Route {
	db := database.SetupDBConnection()
	userRepository := repository.NewUserRepository(db)
	personRepository := repository.NewPersonRepository(db)
	userRolesRepository := repository.NewUserRoleRepository(db)
	sectionsRepository := repository.NewSectionRepository(db)
	userService := services.NewUserService(userRepository)
	jwtService := services.NewJWTService()
	personService := services.NewPersonService(personRepository)
	userRolesService := services.NewUserRolesServices(userRolesRepository)
	sectionsService := services.NewSectionsServices(sectionsRepository)
	authHandler := handler.NewAuthHandler(userService, jwtService, personService, userRolesService, sectionsService)
	return []helper.Route{
		{
			Method:  "POST",
			Path:    "/register",
			Handler: []gin.HandlerFunc{authHandler.Register},
		}, {
			Method:  "POST",
			Path:    "/login",
			Handler: []gin.HandlerFunc{authHandler.Login},
		},
	}
}
