package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sip/simru/entity"
	"github.com/sip/simru/response"
)

type JWTServices interface {
	GenerateToken(user entity.Users) response.ResponseCredential
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secret string
	issuer string
}

type jwtCustomClaim struct {
	UserID   uint64           `json:"user_id"`
	Role     entity.UserRoles `json:"role"`
	Person   entity.Persons   `json:"person"`
	Email    string           `json:"email"`
	Username string           `json:"username"`
	jwt.StandardClaims
}

func NewJWTService() JWTServices {
	return &jwtService{
		issuer: "xjx",
		secret: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET KEY")
	if secretKey == "" {
		secretKey = "supersecret"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(user entity.Users) response.ResponseCredential {
	claims := &jwtCustomClaim{}
	claims.UserID = user.ID
	claims.Person = user.Person
	claims.Role = user.UserRole
	claims.Username = user.Username
	claims.ExpiresAt = time.Now().AddDate(1, 0, 0).Unix()
	claims.Issuer = j.issuer
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		panic(err)
	}
	credential := response.ResponseCredential{}
	credential.UserID = claims.UserID
	credential.Token = t
	credential.Issuer = claims.Issuer
	credential.IssuedAt = claims.IssuedAt
	credential.ExpiresAt = claims.ExpiresAt
	credential.Role = claims.Role
	credential.Person = claims.Person

	return credential
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtString := strings.Split(token, "Bearer ")[1]
	return jwt.Parse(jwtString, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
