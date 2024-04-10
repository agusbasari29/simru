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
	userService services.UserService
	jwtService  services.JWTServices
	// personService services.PersonService
}

func NewAuthHandler(userService services.UserService, jwtService services.JWTServices) *authHandler {
	return &authHandler{userService, jwtService}
}

func (h *authHandler) Register(c *gin.Context) {
	var request request.UserRegisterRequest
	errRequest := c.ShouldBind(&request)
	if errRequest != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValidation := validate.Struct(request)
	if errValidation != nil {
		errorFormatter := helper.ErrorFormatter(errValidation)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	username := entity.Users{Username: request.Username}
	if h.userService.IsExist(username) {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User is already registered!", nil)
		c.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	newUser, err := h.userService.RegisterUser(request)
	if err != nil {
		errorFomater := helper.ErrorFormatter(err)
		errorMessage := helper.M{"error": errorFomater}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	generateToken := h.jwtService.GenerateToken(newUser)
	userData := response.UserResponseFormatter(newUser)
	data := response.UserDataResponseFormatter(userData, generateToken)
	response := helper.ResponseFormatter(http.StatusOK, "success", "User successfully registered", data)
	c.JSON(http.StatusOK, response)
}

func (h *authHandler) Login(c *gin.Context) {
	var request request.AuthLoginRequest
	errRequest := c.ShouldBind(&request)
	if errRequest != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValidation := validate.Struct(request)
	if errValidation != nil {
		errorFormatter := helper.ErrorFormatter(errValidation)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	credential := h.userService.VerifyCredential(request)
	if v, ok := credential.(entity.Users); ok {
		generatedToken := h.jwtService.GenerateToken(v)
		userData := response.UserResponseFormatter(v)
		data := response.UserDataResponseFormatter(userData, generatedToken)
		response := helper.ResponseFormatter(http.StatusOK, "success", "User successfully logedin", data)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusUnauthorized, "error", "User failed to login!", nil)
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
