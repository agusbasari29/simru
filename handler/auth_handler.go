package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sip/simru/entity"
	"github.com/sip/simru/helper"
	"github.com/sip/simru/request"
	"github.com/sip/simru/response"
	"github.com/sip/simru/services"
)

var validate = validator.New()

type authHandler struct {
	userService      services.UserService
	jwtService       services.JWTServices
	personService    services.PersonService
	userRolesService services.UserRolesService
	sectionsService  services.SectionsServices
}

func NewAuthHandler(userService services.UserService, jwtService services.JWTServices, personService services.PersonService, userRolesService services.UserRolesService, sectionsService services.SectionsServices) *authHandler {
	return &authHandler{userService, jwtService, personService, userRolesService, sectionsService}
}

func (h *authHandler) Register(c *gin.Context) {
	var requests request.UserRegisterRequest
	errRequest := c.ShouldBind(&requests)
	if errRequest != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValidation := validate.Struct(requests)
	if errValidation != nil {
		errorFormatter := helper.ErrorFormatter(errValidation)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	username := entity.Users{Username: requests.Username}
	if h.userService.IsExist(username) {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User is already registered!", nil)
		c.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	newPerson, err := h.personService.CreatePerson(requests)
	if err != nil {
		errorFomater := helper.ErrorFormatter(err)
		errorMessage := helper.M{"error": errorFomater}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.RegisterUser(requests, newPerson.ID)
	if err != nil {
		errorFomater := helper.ErrorFormatter(err)
		errorMessage := helper.M{"error": errorFomater}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	getUserRole := request.UserRolesRequest{ID: newUser.UserRoleID}
	userRoles, err := h.userRolesService.GetUserRoles(getUserRole)
	if err != nil {
		panic(err.Error())
	}
	getSection := request.SectionRequest{ID: userRoles.SectionID}
	sections, err := h.sectionsService.GetSections(getSection)
	if err != nil {
		panic(err.Error())
	}
	generateToken := h.jwtService.GenerateToken(newUser)
	userData := response.UserResponseFormatter(newUser, userRoles, newPerson, sections)
	data := response.UserDataResponseFormatter(userData, generateToken)
	response := helper.ResponseFormatter(http.StatusOK, "success", "User successfully registered", data)
	c.JSON(http.StatusOK, response)
}

func (h *authHandler) Login(c *gin.Context) {
	var requests request.AuthLoginRequest
	errRequest := c.ShouldBind(&requests)
	if errRequest != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValidation := validate.Struct(requests)
	if errValidation != nil {
		errorFormatter := helper.ErrorFormatter(errValidation)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	credential := h.userService.VerifyCredential(requests)
	if v, ok := credential.(entity.Users); ok {
		generatedToken := h.jwtService.GenerateToken(v)
		getPerson := request.PersonRequest{ID: v.PersonID}
		person, err := h.personService.GetPerson(getPerson)
		if err != nil {
			panic(err.Error())
		}
		getUserRoles := request.UserRolesRequest{ID: v.UserRoleID}
		userRoles, err := h.userRolesService.GetUserRoles(getUserRoles)
		if err != nil {
			panic(err.Error())
		}
		getSections := request.SectionRequest{ID: userRoles.SectionID}
		sections, err := h.sectionsService.GetSections(getSections)
		if err != nil {
			panic(err.Error())
		}
		userData := response.UserResponseFormatter(v, userRoles, person, sections)
		data := response.UserDataResponseFormatter(userData, generatedToken)
		response := helper.ResponseFormatter(http.StatusOK, "success", "User successfully logedin", data)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusUnauthorized, "error", "User failed to login!", nil)
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
