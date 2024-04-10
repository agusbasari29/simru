package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sip/simru/helper"
	"github.com/sip/simru/request"
	"github.com/sip/simru/services"
)

type userHandler struct {
	userService services.UserService
	jwtService  services.JWTServices
}

func NewUserHandler(userService services.UserService, jwtSevice services.JWTServices) *userHandler {
	return &userHandler{userService, jwtSevice}
}

func (h *userHandler) GetUser(c *gin.Context) {
	var request request.UserRequest
	authHeader := c.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization are needed!", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	request.ID = uint64(id)
	user, err := h.userService.GetUser(request)
	if err != nil {
		panic(err.Error())
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "User get successfully", user)
	c.JSON(http.StatusOK, response)
}
