package response

import (
	"time"

	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID        uint64           `json:"id"`
	Username  string           `json:"username"`
	UserRole  entity.UserRoles `json:"user_role"`
	Person    entity.Persons   `json:"person"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
}

type UserDataResponse struct {
	User       interface{} `json:"user_data"`
	Credential interface{} `json:"credential"`
}

func UserResponseFormatter(user entity.Users) UserResponse {
	formatter := UserResponse{}
	formatter.ID = user.ID
	formatter.Username = user.Username
	formatter.UserRole = user.UserRole
	formatter.Person = user.Person
	formatter.CreatedAt = user.CreatedAt
	formatter.UpdatedAt = user.UpdatedAt
	formatter.DeletedAt = user.DeletedAt

	return formatter
}

func UserDataResponseFormatter(user interface{}, credential interface{}) UserDataResponse {
	userData := UserDataResponse{
		User:       user,
		Credential: credential,
	}

	return userData
}
