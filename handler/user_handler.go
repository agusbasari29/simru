package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sip/simru/helper"
	"github.com/sip/simru/request"
	"github.com/sip/simru/response"
	"github.com/sip/simru/services"
)

type userHandler struct {
	userService      services.UserService
	jwtService       services.JWTServices
	personService    services.PersonService
	userRolesService services.UserRolesService
	sectionsServices services.SectionsServices
}

func NewUserHandler(userService services.UserService, jwtSevice services.JWTServices, personService services.PersonService, userRolesService services.UserRolesService, sectionsServices services.SectionsServices) *userHandler {
	return &userHandler{userService, jwtSevice, personService, userRolesService, sectionsServices}
}

func (h *userHandler) GetUser(c *gin.Context) {
	var userRequest request.UserRequest
	authHeader := c.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization are needed!", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	userRequest.ID = uint64(id)
	user, err := h.userService.GetUser(userRequest)
	if err != nil {
		panic(err.Error())
	}
	getPerson := request.PersonRequest{ID: user.PersonID}
	person, err := h.personService.GetPerson(getPerson)
	if err != nil {
		panic(err.Error())
	}
	getUserRoles := request.UserRolesRequest{ID: user.UserRoleID}
	userRoles, err := h.userRolesService.GetUserRoles(getUserRoles)
	if err != nil {
		panic(err.Error())
	}
	generatedToken := h.jwtService.GenerateToken(user)
	getSections := request.SectionRequest{ID: userRoles.SectionID}
	sections, err := h.sectionsServices.GetSections(getSections)
	if err != nil {
		panic(err.Error())
	}
	userData := response.UserResponseFormatter(user, userRoles, person, sections)
	data := response.UserDataResponseFormatter(userData, generatedToken)
	response := helper.ResponseFormatter(http.StatusOK, "success", "User get successfully", data)
	c.JSON(http.StatusOK, response)
}
